package certificate

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var certificateCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	certificateCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
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
