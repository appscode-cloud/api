package ssh

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var secureShellGetRequestSchema *gojsonschema.Schema

func init() {
	var err error
	secureShellGetRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
