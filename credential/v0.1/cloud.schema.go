package credential

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var cloudCredentialUpdateRequestSchema *gojsonschema.Schema
var cloudCredentialDeleteRequestSchema *gojsonschema.Schema
var cloudCredentialCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	cloudCredentialUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "data":{
      "items":{
        "$ref":"#/definitions/credentialCloudCredentialUpdateRequestDataEntry"
      },
      "type":"array"
    },
    "name":{
      "type":"string"
    },
    "provider":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	cloudCredentialDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	cloudCredentialCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "data":{
      "items":{
        "$ref":"#/definitions/credentialCloudCredentialCreateRequestDataEntry"
      },
      "type":"array"
    },
    "name":{
      "type":"string"
    },
    "provider":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *CloudCredentialUpdateRequest) InValid() (*gojsonschema.Result, error) {
	return cloudCredentialUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CloudCredentialDeleteRequest) InValid() (*gojsonschema.Result, error) {
	return cloudCredentialDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CloudCredentialCreateRequest) InValid() (*gojsonschema.Result, error) {
	return cloudCredentialCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
