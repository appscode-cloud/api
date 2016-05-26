package billing

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var invoiceRequestSchema *gojsonschema.Schema

func init() {
	var err error
	invoiceRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "time_unit": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *InvoiceRequest) IsValid() (*gojsonschema.Result, error) {
	return invoiceRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *InvoiceRequest) IsRequest() {}

func (m *InvoiceResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
