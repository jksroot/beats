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

package santa

import (
	"github.com/jksroot/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "santa", asset.ModuleFieldsPri, AssetSanta); err != nil {
		panic(err)
	}
}

// AssetSanta returns asset data.
// This is the base64 encoded gzipped contents of module/santa.
func AssetSanta() string {
	return "eJyUk82O2jAQgO88xWhP7WFRl4o9cKiUQvqjghaRldpb5Y0nxErsiWynLW9f2UkhJPGWcCJj+/tm7Jl7KPC0AsOUZTMAK2yJK7j7THQsERIXvpsBcDSpFpUVpFbwYQYAzRrsiNclzgAygSU3K790D4pJvFDdz54qXMFRU121kRHmBdN+XlgsdRvP4X/AAk+/SfNOHP8wWbki4h/xuhO/0kUNbWDhmAoz0RNtt0/fQ6JNCwSbM9tcCAdLVMyHco3MTFOv48NzyHzwNMhIg83RVWZ8JiNiSRynaHch51OFmlmhjh4JlDVdMqLkwhQDZbc7BuxPvjN8PZuvybdov4+jQ9sWZt451e3EvvYXlbXEq6We5jnHdpc/0km9y3mpzf8grkK3DypNllIqAyiDWrDyJlqzFVQtX1CHMjPc/enxzk/nOA/m/U26j8nmtTtwT9zP+zIU+/02hiTZQLJ7t3xYbG8yemRAl/Vv/Oxi1WBtHN++ayEUhzeZKNGcjEXpu+9tsMpa2Sn0itl8pOFT1FZkImUW5ylJSepn751CY3dlW/ujngmZJgmpGzQjjspNXVfyegomZ4vl41R78iVaLB8hZyZ3wx12/w0AAP//xDi+7g=="
}
