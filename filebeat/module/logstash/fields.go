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

package logstash

import (
	"github.com/jksroot/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "logstash", asset.ModuleFieldsPri, AssetLogstash); err != nil {
		panic(err)
	}
}

// AssetLogstash returns asset data.
// This is the base64 encoded gzipped contents of module/logstash.
func AssetLogstash() string {
	return "eJzsVU1v2zAMvedXED23+QE+9LKtQIB9ANvuhmLTMhdJFPSRxP9+sB0ntiN3SNsVGDDdLFp875FP1APssMlAsfRB+HoFECgozOBu2LpbAZToC0c2EJsMHlcAcD4BX7iMClcAFaEqfdZFH8AIjZO87QqNxQyk42hPO4nM00yzbOe9C9HPI6Ln2BXQIli/njpIqBxrCDXCkLRTsB79Ouc25qeHUoxXT2WHzYFdOYs9Q6hdP2s85QR2UCjhPRxqdNhRxD2aAOxIkhEB10lKoXYo5rCvoLQxFTst2jCILcfQUXHRGDLyhDbiqFguMQTQUQXKUwWdKMBjuAoOGq6CI6vkXYGS0nn7C4t56A/Kd9iAMCXshYoIJW6jlK1mulQk3QJLFhUZzOmWPuBRaNv6Wwsyt9tm8xG46lowwF/Ipb2L3guZNq9QJPwsYkWol05pkk707IKLmG4Q7lHdiKZYrlPnlvAGLK/4MJsdt86H6xT/B8E/MQiWh8CLZH8XByijtsPtOmlSSZy/qcuqKMnk7cfbqfsqNJ7nRgfwHHYL9IYeb+wMO4ONsTH4e3giFdD5e/gWQ7vT3oIPXGLhF8zOvMvJ5JqUovks6TkqNvI2gp+OWMTO74E0QsVuxBXIQI+GBZtygdepcFY4odO0XlS6H8G1l65/mSYlhIJNRTL24/H97dkrzZPv7eue4ofHqd6JUIgeS9g2o0qkG/IerxAkjGmE4bQBlnG7CbMuF3o5A/8dAAD//7rhE3k="
}
