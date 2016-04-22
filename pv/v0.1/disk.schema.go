package pv

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var diskDescribeRequestSchema *gojsonschema.Schema
var diskDeleteRequestSchema *gojsonschema.Schema
var diskListRequestSchema *gojsonschema.Schema
var diskCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	diskDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cluster":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "plugin":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	diskDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cluster":{
      "type":"string"
    },
    "identifier":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	diskListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cluster":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	diskCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cluster":{
      "type":"string"
    },
    "disk_type":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "size":{
      "type":"integer"
    },
    "zone":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *DiskDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return diskDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DiskDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return diskDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DiskListRequest) IsValid() (*gojsonschema.Result, error) {
	return diskListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DiskCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return diskCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}