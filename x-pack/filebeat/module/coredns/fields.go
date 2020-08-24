// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package coredns

import (
	"github.com/jksroot/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "coredns", asset.ModuleFieldsPri, AssetCoredns); err != nil {
		panic(err)
	}
}

// AssetCoredns returns asset data.
// This is the base64 encoded gzipped contents of module/coredns.
func AssetCoredns() string {
	return "eJy00kFugzAQheE9p3gXSA7Aopt22256gMqxB2LF8VCPUeWcvrIhBCFQpKKyikb4/8YKB1wo1dAcyHipgGijoxqv08CQ6GC7aNnXeKkA4J1N7wgNB5yVN876Fo5bQRfY9JoMTmmWbCw5I3U5eoBXV5qD+Ympoxpt4L4bJytqfsZjYxKqiRTgOVyVszeV3x5fnZsP1ZqpNJAXSj8cHtMNFrAG3CCeCW8fn4hBeVG6eAviu6eQjmJvtKCsj9RSmKZN3jrWOKVI8nyBXJyvUJx1XDsl8teLTm0MmVUh/94PlMpqPyf390tl0Q8kHXuho2azi7iHUEJbSuNUu+uvmJihtOX8//d2p5YrGC9C+osvC/7E7Ej559AQKPerqt8AAAD//23BSVQ="
}
