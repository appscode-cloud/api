package certificate

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var certificateImportRequestSchema *gojsonschema.Schema
var certificateRenewRequestSchema *gojsonschema.Schema
var certificateDeleteRequestSchema *gojsonschema.Schema
var certificateRevokeRequestSchema *gojsonschema.Schema
var certificateDescribeRequestSchema *gojsonschema.Schema
var certificateCreateRequestSchema *gojsonschema.Schema
var certificateDeployRequestSchema *gojsonschema.Schema

func init() {
	var err error
	certificateImportRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	certificateRenewRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "account_phid": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	certificateDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	certificateRevokeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "account_phid": {
      "type": "string"
    },
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	certificateDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	certificateCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "certificateSubjectInfo": {
      "properties": {
        "C": {
          "type": "string"
        },
        "L": {
          "type": "string"
        },
        "O": {
          "type": "string"
        },
        "OU": {
          "type": "string"
        },
        "ST": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
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
    },
    "subject_info": {
      "$ref": "#/definitions/certificateSubjectInfo"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	certificateDeployRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    },
    "secret_name": {
      "type": "string"
    },
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *CertificateImportRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateImportRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateImportRequest) IsRequest() {}

func (m *CertificateRenewRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateRenewRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateRenewRequest) IsRequest() {}

func (m *CertificateDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDeleteRequest) IsRequest() {}

func (m *CertificateRevokeRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateRevokeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateRevokeRequest) IsRequest() {}

func (m *CertificateDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDescribeRequest) IsRequest() {}

func (m *CertificateCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateCreateRequest) IsRequest() {}

func (m *CertificateDeployRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateDeployRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDeployRequest) IsRequest() {}

