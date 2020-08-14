// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package storage

import (
	"os"

	"github.com/JitendraKSahu/beats/v7/metricbeat/mb"

	// Register input module and metricset
	_ "github.com/JitendraKSahu/beats/v7/x-pack/metricbeat/module/googlecloud"
	_ "github.com/JitendraKSahu/beats/v7/x-pack/metricbeat/module/googlecloud/metrics"
)

func init() {
	// To be moved to some kind of helper
	os.Setenv("BEAT_STRICT_PERMS", "false")
	mb.Registry.SetSecondarySource(mb.NewLightModulesSource("../../../module"))
}