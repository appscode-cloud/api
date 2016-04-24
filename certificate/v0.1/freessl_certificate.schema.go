package certificate

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var freeSSLCertificateRevokeRequestSchema *gojsonschema.Schema
var freeSSLCertificateRegisterRequestSchema *gojsonschema.Schema
var freeSSLCertificateRenewRequestSchema *gojsonschema.Schema
var freeSSLCertificateCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	freeSSLCertificateRevokeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "account_phid": {
      "type": "string"
    },
    "cert_phid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	freeSSLCertificateRegisterRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "email": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	freeSSLCertificateRenewRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "account_phid": {
      "type": "string"
    },
    "cert_phid": {
      "type": "string"
    },
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	freeSSLCertificateCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "account_phid": {
      "type": "string"
    },
    "bundle": {
      "type": "boolean"
    },
    "common_name": {
      "type": "string"
    },
    "key_data": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "san": {
      "items": {
        "type": "string"
      },
      "type": "array"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *FreeSSLCertificateRevokeRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateRevokeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateRevokeRequest) IsRequest() {}

func (m *FreeSSLCertificateRegisterRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateRegisterRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateRegisterRequest) IsRequest() {}

func (m *FreeSSLCertificateRenewRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateRenewRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateRenewRequest) IsRequest() {}

func (m *FreeSSLCertificateCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateCreateRequest) IsRequest() {}

