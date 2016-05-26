package billing

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var purchaseConfirmRequestSchema *gojsonschema.Schema
var purchaseBeginRequestSchema *gojsonschema.Schema
var purchaseCloseRequestSchema *gojsonschema.Schema

func init() {
	var err error
	purchaseConfirmRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "object_phid": {
      "type": "string"
    },
    "phid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	purchaseBeginRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "count": {
      "type": "integer"
    },
    "product_sku": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	purchaseCloseRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "object_phid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *PurchaseConfirmRequest) IsValid() (*gojsonschema.Result, error) {
	return purchaseConfirmRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PurchaseConfirmRequest) IsRequest() {}

func (m *PurchaseBeginRequest) IsValid() (*gojsonschema.Result, error) {
	return purchaseBeginRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PurchaseBeginRequest) IsRequest() {}

func (m *PurchaseCloseRequest) IsValid() (*gojsonschema.Result, error) {
	return purchaseCloseRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PurchaseCloseRequest) IsRequest() {}

func (m *PurchaseBeginResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
