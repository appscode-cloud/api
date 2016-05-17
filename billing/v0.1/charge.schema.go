package billing

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var chargeRequestSchema *gojsonschema.Schema

func init() {
	var err error
	chargeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ChargeRequest) IsValid() (*gojsonschema.Result, error) {
	return chargeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ChargeRequest) IsRequest() {}

