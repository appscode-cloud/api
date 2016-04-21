package ci

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var slaveDeleteRequestSchema *gojsonschema.Schema
var slaveRestartRequestSchema *gojsonschema.Schema
var slaveCreateRequestSchema *gojsonschema.Schema
var slaveDescribeRequestSchema *gojsonschema.Schema

func init() {
	var err error
	slaveDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	slaveRestartRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	slaveCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "ci_starter_version":{
      "type":"string"
    },
    "executors":{
      "type":"integer"
    },
    "labels":{
      "type":"string"
    },
    "ports":{
      "items":{
        "$ref":"#/definitions/ciPortInfo"
      },
      "type":"array"
    },
    "saltbase_version":{
      "type":"string"
    },
    "user_startup_script":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	slaveDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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

func (m *SlaveDeleteRequest) InValid() (*gojsonschema.Result, error) {
	return slaveDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SlaveRestartRequest) InValid() (*gojsonschema.Result, error) {
	return slaveRestartRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SlaveCreateRequest) InValid() (*gojsonschema.Result, error) {
	return slaveCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SlaveDescribeRequest) InValid() (*gojsonschema.Result, error) {
	return slaveDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
