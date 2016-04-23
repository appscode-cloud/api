package credential

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var cloudCredentialUpdateRequestSchema *gojsonschema.Schema
var cloudCredentialDeleteRequestSchema *gojsonschema.Schema
var cloudCredentialCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	cloudCredentialUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "credentialCloudCredentialUpdateRequestDataEntry": {
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "data": {
      "items": {
        "$ref": "#/definitions/credentialCloudCredentialUpdateRequestDataEntry"
      },
      "type": "array"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
      "type": "string"
    },
    "provider": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	cloudCredentialDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	cloudCredentialCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "credentialCloudCredentialCreateRequestDataEntry": {
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "data": {
      "items": {
        "$ref": "#/definitions/credentialCloudCredentialCreateRequestDataEntry"
      },
      "type": "array"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
      "type": "string"
    },
    "provider": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *CloudCredentialUpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return cloudCredentialUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CloudCredentialUpdateRequest) IsRequest() {}

func (m *CloudCredentialDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return cloudCredentialDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CloudCredentialDeleteRequest) IsRequest() {}

func (m *CloudCredentialCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return cloudCredentialCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CloudCredentialCreateRequest) IsRequest() {}
