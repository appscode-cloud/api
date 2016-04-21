package ca

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var certificateCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	certificateCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "csr":{
      "type":"string"
    },
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

func (m *CertificateCreateRequest) InValid() (*gojsonschema.Result, error) {
	return certificateCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
