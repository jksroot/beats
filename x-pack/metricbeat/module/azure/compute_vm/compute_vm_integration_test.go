// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build integration
// +build azure

package compute_vm

import (
	"testing"

	"github.com/JitendraKSahu/beats/v7/x-pack/metricbeat/module/azure/test"

	"github.com/stretchr/testify/assert"

	mbtest "github.com/JitendraKSahu/beats/v7/metricbeat/mb/testing"
)

func TestFetchMetricset(t *testing.T) {
	config := test.GetConfig(t, "compute_vm")
	metricSet := mbtest.NewReportingMetricSetV2Error(t, config)
	events, errs := mbtest.ReportingFetchV2Error(metricSet)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)
	mbtest.TestMetricsetFieldsDocumented(t, metricSet, events)
}

func TestData(t *testing.T) {
	config := test.GetConfig(t, "compute_vm")
	metricSet := mbtest.NewFetcher(t, config)
	metricSet.WriteEvents(t, "/")
}