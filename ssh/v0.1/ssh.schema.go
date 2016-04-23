package ssh

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var secureShellGetRequestSchema *gojsonschema.Schema

func init() {
	var err error
	secureShellGetRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
      "type": "string"
    },
    "cluster_namespace": {
      "type": "string"
    },
    "instance_name": {
      "type": "string"
    },
    "jenkins_namespace": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *SecureShellGetRequest) IsValid() (*gojsonschema.Result, error) {
	return secureShellGetRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SecureShellGetRequest) IsRequest() {}
