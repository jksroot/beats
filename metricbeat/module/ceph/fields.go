// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package ceph

import (
	"github.com/jksroot/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "ceph", asset.ModuleFieldsPri, AssetCeph); err != nil {
		panic(err)
	}
}

// AssetCeph returns asset data.
// This is the base64 encoded gzipped contents of module/ceph.
func AssetCeph() string {
	return "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"
}
