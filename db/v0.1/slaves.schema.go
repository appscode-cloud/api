package db

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var slaveAddRequestSchema *gojsonschema.Schema

func init() {
	var err error
	slaveAddRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "config": {
      "type": "string"
    },
    "name": {
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

func (m *SlaveAddRequest) IsValid() (*gojsonschema.Result, error) {
	return slaveAddRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SlaveAddRequest) IsRequest() {}

