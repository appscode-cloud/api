package billing

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var chargeCalculateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	chargeCalculateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "charge_type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ChargeCalculateRequest) IsValid() (*gojsonschema.Result, error) {
	return chargeCalculateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ChargeCalculateRequest) IsRequest() {}

