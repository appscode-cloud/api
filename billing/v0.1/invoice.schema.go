package billing

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

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
		glog.Fatal(err)
	}
}

func (m *ChargeRequest) IsValid() (*gojsonschema.Result, error) {
	return chargeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ChargeRequest) IsRequest() {}

func (m *ChargeResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
