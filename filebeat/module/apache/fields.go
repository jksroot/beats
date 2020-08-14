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

package apache

import (
	"github.com/JitendraKSahu/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "apache", asset.ModuleFieldsPri, AssetApache); err != nil {
		panic(err)
	}
}

// AssetApache returns asset data.
// This is the base64 encoded gzipped contents of module/apache.
func AssetApache() string {
	return "eJysl7+O4zgMxvs8BTH9qLgyxQGHAw5X3AELzPSGIjO2MIropegM/PYL/5tNMrIjOXEVyOH3/SjSgvgKH9jtQTfa1LgDECsO9/Dy17DwsgMoMRi2jVjye/hzBwAwvoT/qWxdHxRqYikM+aOt9iDc9otHi64M+yHgFbw+4Wzzx7AGIF2De6iY2mZaiXgNfs7qgAGOxHDQ5uNTcwmGTo0We7DOSgefVmogV84Wk/0kcclyxWMMhvC1HEOKhV9KMJ5IsLDN1dtZSffkN28aLfUeArVsUOmy5GuE/jnZivW4C9NufjcOwamGSciQy/Qe90iN2asVoVQOY5sa+WGKqEwKw4HKrgjoRR06wVu7eyC1SKMYQ0M+oOq1ojIpIG1ALvqfmQh9nIrEpXieUGoqt+X8s8UgKqqQlC7ntl7LThHbynq9pdl67OKMHCz5LRnHQ1Oc5/4oDJW51b1usCBa2hDTSeM4InP2x3ZV7wWNFHtdoZcNzV0MgamlX/62lv1vD21YOLgvJUs8W3NbhfWUommNOrEPOJbdEs1CeCbLgxAUM8lEoKCOrXOxgzAPZUkhn+dRlHjfbmFZVkr5/iqkhavGlv435MV69PLIRk83mQpJ3dVL3W9DrRfuChsodlBuQrurmArnyAx/exxqRSkVhrGy5J9Uv3Wx5OJZ6Z7VUCtSmTv0vFa6L7iENiMhM/FDY4fDM+bevRxVKhaXdskMQVe5F594VIpfY3NvtA3TMEh8j0zxk81+UjPqUm1yPc0TdI7xNDcNPaSiCt+uU5EhPHsGH2f+sSfVc+bqBav++Zu8aOvDZDGM/lLjjPHv+/sPeEM+I09mfXd/ccXYIGd2/sDuk/i2pCu8/fP29h/MqjBNGr+JtgzPmzFGzcFEbTx9Vlz+mWrCdLosyqCbVYeV/t+U+nuNk2ZfiLI11lcDoaOqwnI+j9TuVwAAAP//QYrRgg=="
}