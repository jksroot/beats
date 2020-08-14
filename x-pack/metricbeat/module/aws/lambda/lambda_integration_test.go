// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build integration
// +build aws

package lambda

import (
	"testing"

	mbtest "github.com/JitendraKSahu/beats/v7/metricbeat/mb/testing"
	"github.com/JitendraKSahu/beats/v7/x-pack/metricbeat/module/aws/mtest"
)

func TestData(t *testing.T) {
	config := mtest.GetConfigForTest(t, "lambda", "300s")

	metricSet := mbtest.NewFetcher(t, config)
	metricSet.WriteEvents(t, "/")
}