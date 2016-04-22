package namespace

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var statusRequestSchema *gojsonschema.Schema
var createRequestSchema *gojsonschema.Schema
var checkRequestSchema *gojsonschema.Schema
var logRequestSchema *gojsonschema.Schema

func init() {
	var err error
	statusRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	createRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "display_name":{
      "type":"string"
    },
    "email":{
      "type":"string"
    },
    "invite_email":{
      "items":{
        "type":"string"
      },
      "type":"array"
    },
    "name":{
      "type":"string"
    },
    "password":{
      "type":"string"
    },
    "subscription_type":{
      "type":"string"
    },
    "user_ip":{
      "type":"string"
    },
    "user_name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	checkRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	logRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *StatusRequest) IsValid() (*gojsonschema.Result, error) {
	return statusRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CreateRequest) IsValid() (*gojsonschema.Result, error) {
	return createRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CheckRequest) IsValid() (*gojsonschema.Result, error) {
	return checkRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *LogRequest) IsValid() (*gojsonschema.Result, error) {
	return logRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
