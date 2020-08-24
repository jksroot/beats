// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package s3

import (
	"github.com/jksroot/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "s3", asset.ModuleFieldsPri, AssetS3); err != nil {
		panic(err)
	}
}

// AssetS3 returns asset data.
// This is the base64 encoded gzipped contents of input/s3.
func AssetS3() string {
	return "eJykjjGugzAQRHufYkQPjTsX/wi/+QdABg8fB4ORvSTi9pGBJlKkFNlipZ3dnTc1Ju4GWStAvAQaVFlXCnDMffKr+LgY/CgA+NMYPIPLGFKckTX8sm7SKCAx0GYadBSrcN2Z463GYuey2vqJ0pbh0AHZV5qS4BGTu7Q33FK/dibiABlZcpxekNGW5jNC/EeiJM873ZGveYHH7sZe2on71+zT6gP7GQAA//+k2GkG"
}
