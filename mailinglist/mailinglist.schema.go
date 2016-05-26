package mailinglist

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var subscribeRequestSchema *gojsonschema.Schema
var unsubscribeRequestSchema *gojsonschema.Schema
var sendEmailRequestSchema *gojsonschema.Schema

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
		glog.Fatal(err)
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
		glog.Fatal(err)
	}
	sendEmailRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "body": {
      "type": "string"
    },
    "sender_email": {
      "type": "string"
    },
    "sender_name": {
      "type": "string"
    },
    "subject": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
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

func (m *SendEmailRequest) IsValid() (*gojsonschema.Result, error) {
	return sendEmailRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SendEmailRequest) IsRequest() {}
