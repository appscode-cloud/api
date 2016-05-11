package mailinglist

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var subscribeRequestSchema *gojsonschema.Schema
var unsubscribeRequestSchema *gojsonschema.Schema

func init() {
	var err error
	subscribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	unsubscribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "magic_code": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *SubscribeRequest) IsValid() (*gojsonschema.Result, error) {
	return subscribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscribeRequest) IsRequest() {}

func (m *UnsubscribeRequest) IsValid() (*gojsonschema.Result, error) {
	return unsubscribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *UnsubscribeRequest) IsRequest() {}

