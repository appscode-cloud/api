package certificate

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var freeSSLCertificateRevokeRequestSchema *gojsonschema.Schema
var freeSSLCertificateDeployRequestSchema *gojsonschema.Schema
var freeSSLCertificateRegisterRequestSchema *gojsonschema.Schema
var freeSSLCertificateCreateRequestSchema *gojsonschema.Schema
var freeSSLCertificateDeleteRequestSchema *gojsonschema.Schema
var freeSSLCertificateDescribeRequestSchema *gojsonschema.Schema
var freeSSLCertificateRenewRequestSchema *gojsonschema.Schema

func init() {
	var err error
	freeSSLCertificateRevokeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cert": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	freeSSLCertificateDeployRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cert": {
      "type": "string"
    },
    "cluster_name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    },
    "secret_name": {
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
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	freeSSLCertificateCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cert_data": {
      "type": "string"
    },
    "key_data": {
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
	freeSSLCertificateDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cert": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	freeSSLCertificateDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cert": {
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
    "cert": {
      "type": "string"
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

func (m *FreeSSLCertificateDeployRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateDeployRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateDeployRequest) IsRequest() {}

func (m *FreeSSLCertificateRegisterRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateRegisterRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateRegisterRequest) IsRequest() {}

func (m *FreeSSLCertificateCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateCreateRequest) IsRequest() {}

func (m *FreeSSLCertificateDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateDeleteRequest) IsRequest() {}

func (m *FreeSSLCertificateDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateDescribeRequest) IsRequest() {}

func (m *FreeSSLCertificateRenewRequest) IsValid() (*gojsonschema.Result, error) {
	return freeSSLCertificateRenewRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *FreeSSLCertificateRenewRequest) IsRequest() {}

