package backup

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var clientReConfigureRequestSchema *gojsonschema.Schema
var diskAddRequestSchema *gojsonschema.Schema

func init() {
	var err error
	clientReConfigureRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{3,61}[a-z0-9])?$",
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{3,61}[a-z0-9])?$",
      "type": "string"
    },
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{3,61}[a-z0-9])?$",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	diskAddRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{3,61}[a-z0-9])?$",
      "type": "string"
    },
    "disk_name": {
      "type": "string"
    },
    "disk_type": {
      "type": "string"
    },
    "mount_path": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{3,61}[a-z0-9])?$",
      "type": "string"
    },
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{3,61}[a-z0-9])?$",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *ClientReConfigureRequest) IsValid() (*gojsonschema.Result, error) {
	return clientReConfigureRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClientReConfigureRequest) IsRequest() {}

func (m *DiskAddRequest) IsValid() (*gojsonschema.Result, error) {
	return diskAddRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DiskAddRequest) IsRequest() {}

