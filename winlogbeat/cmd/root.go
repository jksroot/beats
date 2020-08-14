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

package cmd

import (
	"github.com/JitendraKSahu/beats/v7/libbeat/cmd"
	"github.com/JitendraKSahu/beats/v7/libbeat/cmd/instance"
	"github.com/JitendraKSahu/beats/v7/libbeat/common"
	"github.com/JitendraKSahu/beats/v7/libbeat/publisher/processing"
	"github.com/JitendraKSahu/beats/v7/winlogbeat/beater"

	// Register fields.
	_ "github.com/JitendraKSahu/beats/v7/winlogbeat/include"

	// Import processors and supporting modules.
	_ "github.com/JitendraKSahu/beats/v7/libbeat/processors/script"
	_ "github.com/JitendraKSahu/beats/v7/libbeat/processors/timestamp"
	_ "github.com/JitendraKSahu/beats/v7/winlogbeat/processors/script/javascript/module/winlogbeat"
)

const (
	// Name of this beat.
	Name = "winlogbeat"

	// ecsVersion specifies the version of ECS that Winlogbeat is implementing.
	ecsVersion = "1.5.0"
)

// withECSVersion is a modifier that adds ecs.version to events.
var withECSVersion = processing.WithFields(common.MapStr{
	"ecs": common.MapStr{
		"version": ecsVersion,
	},
})

// RootCmd to handle beats CLI.
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{
	Name:          Name,
	HasDashboards: true,
	Processing:    processing.MakeDefaultSupport(true, withECSVersion, processing.WithAgentMeta()),
})