package billing

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var purchaseOpenRequestSchema *gojsonschema.Schema
var purchaseCloseRequestSchema *gojsonschema.Schema

func init() {
	var err error
	purchaseOpenRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "author_phid": {
      "type": "string"
    },
    "metadata": {
      "type": "string"
    },
    "object_phid": {
      "type": "string"
    },
    "product_phid": {
      "type": "string"
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
}

func (m *PurchaseOpenRequest) IsValid() (*gojsonschema.Result, error) {
	return purchaseOpenRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PurchaseOpenRequest) IsRequest() {}

func (m *PurchaseCloseRequest) IsValid() (*gojsonschema.Result, error) {
	return purchaseCloseRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PurchaseCloseRequest) IsRequest() {}

