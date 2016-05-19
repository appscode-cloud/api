package namespace

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/xeipuuv/gojsonschema"
	"log"
)
var statusRequestSchema *gojsonschema.Schema
var createRequestSchema *gojsonschema.Schema
var checkRequestSchema *gojsonschema.Schema
var logRequestSchema *gojsonschema.Schema

func init() {
	var err error
	statusRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	createRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "display_name": {
      "type": "string"
    },
    "email": {
      "type": "string"
    },
    "invite_email": {
      "items": {
        "type": "string"
      },
      "type": "array"
    },
    "name": {
      "type": "string"
    },
    "password": {
      "type": "string"
    },
    "subscription_type": {
      "type": "string"
    },
    "user_ip": {
      "type": "string"
    },
    "user_name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	checkRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	logRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *StatusRequest) IsValid() (*gojsonschema.Result, error) {
	return statusRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *StatusRequest) IsRequest() {}

func (m *CreateRequest) IsValid() (*gojsonschema.Result, error) {
	return createRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CreateRequest) IsRequest() {}

func (m *CheckRequest) IsValid() (*gojsonschema.Result, error) {
	return checkRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CheckRequest) IsRequest() {}

func (m *LogRequest) IsValid() (*gojsonschema.Result, error) {
	return logRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *LogRequest) IsRequest() {}

func (m *CheckResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

func (m *StatusResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

func (m *LogResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

func (m *CreateResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

