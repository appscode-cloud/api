package billing

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var purchaseOpenRequestSchema *gojsonschema.Schema
var purchaseQoutaRequestSchema *gojsonschema.Schema
var purchaseCloseRequestSchema *gojsonschema.Schema

func init() {
	var err error
	purchaseOpenRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	purchaseQoutaRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "count": {
      "type": "integer"
    },
    "product_type": {
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

func (m *PurchaseOpenRequest) IsValid() (*gojsonschema.Result, error) {
	return purchaseOpenRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PurchaseOpenRequest) IsRequest() {}

func (m *PurchaseQoutaRequest) IsValid() (*gojsonschema.Result, error) {
	return purchaseQoutaRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PurchaseQoutaRequest) IsRequest() {}

func (m *PurchaseCloseRequest) IsValid() (*gojsonschema.Result, error) {
	return purchaseCloseRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PurchaseCloseRequest) IsRequest() {}

func (m *PurchaseQutaResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
