package mailinglist

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var membershipRequestSchema *gojsonschema.Schema

func init() {
	var err error
	membershipRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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

func (m *MembershipRequest) IsValid() (*gojsonschema.Result, error) {
	return membershipRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *MembershipRequest) IsRequest() {}

