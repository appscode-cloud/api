package authentication

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var validateRequestSchema *gojsonschema.Schema

func init() {
	var err error
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

func (m *ValidateRequest) IsValid() (*gojsonschema.Result, error) {
	return validateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ValidateRequest) IsRequest() {}

func (m *ValidateResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
