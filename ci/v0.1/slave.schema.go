package ci

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var slaveDeleteRequestSchema *gojsonschema.Schema
var slaveRestartRequestSchema *gojsonschema.Schema
var slaveDescribeRequestSchema *gojsonschema.Schema
var slaveCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	slaveDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
		glog.Fatal(err)
	}
	slaveRestartRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
		glog.Fatal(err)
	}
	slaveDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
		glog.Fatal(err)
	}
	slaveCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "ciPortInfo": {
      "properties": {
        "port_range": {
          "type": "string"
        },
        "protocol": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "ci_starter_version": {
      "type": "string"
    },
    "executors": {
      "type": "integer"
    },
    "labels": {
      "type": "string"
    },
    "ports": {
      "items": {
        "$ref": "#/definitions/ciPortInfo"
      },
      "type": "array"
    },
    "saltbase_version": {
      "type": "string"
    },
    "sku": {
      "type": "string"
    },
    "user_startup_script": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *SlaveDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return slaveDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SlaveDeleteRequest) IsRequest() {}

func (m *SlaveRestartRequest) IsValid() (*gojsonschema.Result, error) {
	return slaveRestartRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SlaveRestartRequest) IsRequest() {}

func (m *SlaveDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return slaveDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SlaveDescribeRequest) IsRequest() {}

func (m *SlaveCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return slaveCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SlaveCreateRequest) IsRequest() {}

func (m *SlaveListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *SlaveRestartResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *SlaveDescribeResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
