package certificate

import (
	"github.com/xeipuuv/gojsonschema"
	// "log"
)

// Auto-generated. DO NOT EDIT.

var tLSCertificateCreateRequestSchema *gojsonschema.Schema
var tLSCertificateDeployRequestSchema *gojsonschema.Schema
var tLSCertificateDescribeRequestSchema *gojsonschema.Schema
var tLSCertificateDeleteRequestSchema *gojsonschema.Schema

func init() {
	var err error
	tLSCertificateCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cert_data":{
      "type":"string"
    },
    "key_data":{
      "type":"string"
    },
    "name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		// log.Fatal(err)
	}
	tLSCertificateDeployRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cert":{
      "type":"string"
    },
    "cluster_name":{
      "type":"string"
    },
    "namespace":{
      "type":"string"
    },
    "secret_name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		// log.Fatal(err)
	}
	tLSCertificateDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cert":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		// log.Fatal(err)
	}
	tLSCertificateDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cert":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		// log.Fatal(err)
	}
}

func (m *TLSCertificateCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return tLSCertificateCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *TLSCertificateDeployRequest) IsValid() (*gojsonschema.Result, error) {
	return tLSCertificateDeployRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *TLSCertificateDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return tLSCertificateDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *TLSCertificateDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return tLSCertificateDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
