// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package googlecloud

import (
	"github.com/jksroot/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "googlecloud", asset.ModuleFieldsPri, AssetGooglecloud); err != nil {
		panic(err)
	}
}

// AssetGooglecloud returns asset data.
// This is the base64 encoded gzipped contents of module/googlecloud.
func AssetGooglecloud() string {
	return "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"
}
