package ci

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var buildListRequestSchema *gojsonschema.Schema
var buildDescribeRequestSchema *gojsonschema.Schema

func init() {
	var err error
	buildListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "job_name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	buildDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "console": {
      "type": "string"
    },
    "job_name": {
      "type": "string"
    },
    "number": {
      "type": "integer"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *BuildListRequest) IsValid() (*gojsonschema.Result, error) {
	return buildListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *BuildDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return buildDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
