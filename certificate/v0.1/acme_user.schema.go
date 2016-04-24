package certificate

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var acmeUserRegisterRequestSchema *gojsonschema.Schema

func init() {
	var err error
	acmeUserRegisterRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "email": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *AcmeUserRegisterRequest) IsValid() (*gojsonschema.Result, error) {
	return acmeUserRegisterRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AcmeUserRegisterRequest) IsRequest() {}

