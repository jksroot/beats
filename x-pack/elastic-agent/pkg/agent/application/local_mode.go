// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package application

import (
	"context"

	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/agent/application/filters"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/agent/application/info"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/agent/configrequest"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/agent/configuration"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/agent/errors"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/agent/operation"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/config"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/core/logger"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/core/monitoring"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/core/server"
	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/dir"
	reporting "github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/reporter"
	logreporter "github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/reporter/log"
)

type emitterFunc func(*config.Config) error

// ConfigHandler is capable of handling config, perform actions at it, shutdown any long running process.
type ConfigHandler interface {
	HandleConfig(configrequest.Request) error
	Close() error
	Shutdown()
}

type discoverFunc func() ([]string, error)

// ErrNoConfiguration is returned when no configuration are found.
var ErrNoConfiguration = errors.New("no configuration found", errors.TypeConfig)

// Local represents a standalone agents, that will read his configuration directly from disk.
// Some part of the configuration can be reloaded.
type Local struct {
	bgContext   context.Context
	cancelCtxFn context.CancelFunc
	log         *logger.Logger
	router      *router
	source      source
	agentInfo   *info.AgentInfo
	srv         *server.Server
}

type source interface {
	Start() error
	Stop() error
}

// newLocal return a agent managed by local configuration.
func newLocal(
	ctx context.Context,
	log *logger.Logger,
	pathConfigFile string,
	rawConfig *config.Config,
) (*Local, error) {
	cfg, err := configuration.NewFromConfig(rawConfig)
	if err != nil {
		return nil, err
	}

	if log == nil {
		log, err = logger.NewFromConfig("", cfg.Settings.LoggingConfig)
		if err != nil {
			return nil, err
		}
	}
	agentInfo, err := info.NewAgentInfo()
	if err != nil {
		return nil, err
	}

	logR := logreporter.NewReporter(log)

	localApplication := &Local{
		log:       log,
		agentInfo: agentInfo,
	}

	localApplication.bgContext, localApplication.cancelCtxFn = context.WithCancel(ctx)
	localApplication.srv, err = server.NewFromConfig(log, cfg.Settings.GRPC, &operation.ApplicationStatusHandler{})
	if err != nil {
		return nil, errors.New(err, "initialize GRPC listener")
	}

	reporter := reporting.NewReporter(localApplication.bgContext, log, localApplication.agentInfo, logR)

	monitor, err := monitoring.NewMonitor(cfg.Settings)
	if err != nil {
		return nil, errors.New(err, "failed to initialize monitoring")
	}

	router, err := newRouter(log, streamFactory(localApplication.bgContext, cfg.Settings, localApplication.srv, reporter, monitor))
	if err != nil {
		return nil, errors.New(err, "fail to initialize pipeline router")
	}
	localApplication.router = router

	discover := discoverer(pathConfigFile, cfg.Settings.Path)
	emit := emitter(
		log,
		router,
		&configModifiers{
			Decorators: []decoratorFunc{injectMonitoring},
			Filters:    []filterFunc{filters.StreamChecker, filters.ConstraintFilter},
		},
		monitor,
	)

	var cfgSource source
	if !cfg.Settings.Reload.Enabled {
		log.Debug("Reloading of configuration is off")
		cfgSource = newOnce(log, discover, emit)
	} else {
		log.Debugf("Reloading of configuration is on, frequency is set to %s", cfg.Settings.Reload.Period)
		cfgSource = newPeriodic(log, cfg.Settings.Reload.Period, discover, emit)
	}

	localApplication.source = cfgSource

	return localApplication, nil
}

// Start starts a local agent.
func (l *Local) Start() error {
	l.log.Info("Agent is starting")
	defer l.log.Info("Agent is stopped")

	if err := l.srv.Start(); err != nil {
		return err
	}
	if err := l.source.Start(); err != nil {
		return err
	}

	return nil
}

// Stop stops a local agent.
func (l *Local) Stop() error {
	err := l.source.Stop()
	l.cancelCtxFn()
	l.router.Shutdown()
	l.srv.Stop()
	return err
}

// AgentInfo retrieves agent information.
func (l *Local) AgentInfo() *info.AgentInfo {
	return l.agentInfo
}

func discoverer(patterns ...string) discoverFunc {
	var p []string
	for _, newP := range patterns {
		if len(newP) == 0 {
			continue
		}

		p = append(p, newP)
	}

	if len(p) == 0 {
		return func() ([]string, error) {
			return []string{}, ErrNoConfiguration
		}
	}

	return func() ([]string, error) {
		return dir.DiscoverFiles(p...)
	}
}
