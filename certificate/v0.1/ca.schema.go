package certificate

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var cACreateRequestSchema *gojsonschema.Schema
var cARegisterRequestSchema *gojsonschema.Schema

func init() {
	var err error
	cACreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "csr": {
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
	cARegisterRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "email": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *CACreateRequest) IsValid() (*gojsonschema.Result, error) {
	return cACreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CACreateRequest) IsRequest() {}

func (m *CARegisterRequest) IsValid() (*gojsonschema.Result, error) {
	return cARegisterRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CARegisterRequest) IsRequest() {}

