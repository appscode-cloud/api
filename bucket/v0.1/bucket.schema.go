package bucket

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var bucketListRequestSchema *gojsonschema.Schema

func init() {
	var err error
	bucketListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cloud_credential": {
      "maxLength": 63,
      "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *BucketListRequest) IsValid() (*gojsonschema.Result, error) {
	return bucketListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *BucketListRequest) IsRequest() {}
