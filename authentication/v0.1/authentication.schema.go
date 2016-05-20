package authentication

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var loginRequestSchema *gojsonschema.Schema
var logoutRequestSchema *gojsonschema.Schema
var validateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	loginRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "issue_token": {
      "type": "boolean"
    },
    "namespace": {
      "type": "string"
    },
    "secret": {
      "type": "string"
    },
    "username": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	logoutRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "namespace": {
      "type": "string"
    },
    "token": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	validateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "namespace": {
      "type": "string"
    },
    "secret": {
      "type": "string"
    },
    "username": {
      "type": "string"
    }
  },
  "title": "Next Id 4",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *LoginRequest) IsValid() (*gojsonschema.Result, error) {
	return loginRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *LoginRequest) IsRequest() {}

func (m *LogoutRequest) IsValid() (*gojsonschema.Result, error) {
	return logoutRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *LogoutRequest) IsRequest() {}

func (m *ValidateRequest) IsValid() (*gojsonschema.Result, error) {
	return validateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ValidateRequest) IsRequest() {}

func (m *ValidateResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *LogoutResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *LoginResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
