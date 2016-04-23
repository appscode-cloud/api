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
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "csr": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *CertificateCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateCreateRequest) IsRequest() {}
