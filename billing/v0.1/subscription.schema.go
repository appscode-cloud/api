package billing

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var subscriptionQoutaRequestSchema *gojsonschema.Schema
var subscriptionOpenRequestSchema *gojsonschema.Schema
var subscriptionCreateRequestSchema *gojsonschema.Schema
var subscriptionDescribeRequestSchema *gojsonschema.Schema
var subscriptionCloseRequestSchema *gojsonschema.Schema

func init() {
	var err error
	subscriptionQoutaRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "billingResource": {
      "default": "USER",
      "enum": [
        "USER",
        "CLUSTER",
        "NODE",
        "DB",
        "CI"
      ],
      "type": "string"
    }
  },
  "properties": {
    "count": {
      "type": "integer"
    },
    "object_phid": {
      "type": "string"
    },
    "resource": {
      "$ref": "#/definitions/billingResource"
    },
    "subresource": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	subscriptionOpenRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
		log.Fatal(err)
	}
	subscriptionCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "phid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	subscriptionDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	subscriptionCloseRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
		log.Fatal(err)
	}
}

func (m *SubscriptionQoutaRequest) IsValid() (*gojsonschema.Result, error) {
	return subscriptionQoutaRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionQoutaRequest) IsRequest() {}

func (m *SubscriptionOpenRequest) IsValid() (*gojsonschema.Result, error) {
	return subscriptionOpenRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionOpenRequest) IsRequest() {}

func (m *SubscriptionCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return subscriptionCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionCreateRequest) IsRequest() {}

func (m *SubscriptionDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return subscriptionDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionDescribeRequest) IsRequest() {}

func (m *SubscriptionCloseRequest) IsValid() (*gojsonschema.Result, error) {
	return subscriptionCloseRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionCloseRequest) IsRequest() {}

