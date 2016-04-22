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
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cluster_name":{
      "type":"string"
    },
    "cluster_namespace":{
      "type":"string"
    },
    "instance_name":{
      "type":"string"
    },
    "jenkins_namespace":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *SecureShellGetRequest) InValid() (*gojsonschema.Result, error) {
	return secureShellGetRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
