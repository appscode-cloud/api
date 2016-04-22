package pv

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var pVUnregisterRequestSchema *gojsonschema.Schema
var pVDescribeRequestSchema *gojsonschema.Schema
var pVRegisterRequestSchema *gojsonschema.Schema

func init() {
	var err error
	pVUnregisterRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	pVDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	pVRegisterRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "endpoint": {
      "type": "string"
    },
    "identifier": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "plugin": {
      "type": "string"
    },
    "size": {
      "type": "integer"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *PVUnregisterRequest) IsValid() (*gojsonschema.Result, error) {
	return pVUnregisterRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PVDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return pVDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PVRegisterRequest) IsValid() (*gojsonschema.Result, error) {
	return pVRegisterRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
