package kubernetes

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var secretDescribeRequestSchema *gojsonschema.Schema
var configMapDescribeRequestSchema *gojsonschema.Schema
var appDescribeRequestSchema *gojsonschema.Schema
var clientRequestSchema *gojsonschema.Schema

func init() {
	var err error
	secretDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cluster":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "namespace":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	configMapDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cluster":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "namespace":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	appDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cluster":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "namespace":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	clientRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cluster":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *SecretDescribeRequest) InValid() (*gojsonschema.Result, error) {
	return secretDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ConfigMapDescribeRequest) InValid() (*gojsonschema.Result, error) {
	return configMapDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AppDescribeRequest) InValid() (*gojsonschema.Result, error) {
	return appDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClientRequest) InValid() (*gojsonschema.Result, error) {
	return clientRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
