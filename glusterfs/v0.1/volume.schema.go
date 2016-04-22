package glusterfs

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var volumeListRequestSchema *gojsonschema.Schema
var volumeCreateRequestSchema *gojsonschema.Schema
var volumeDeleteRequestSchema *gojsonschema.Schema

func init() {
	var err error
	volumeListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "glusterfs_cluster":{
      "type":"string"
    },
    "kube_cluster":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	volumeCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "glusterfs_cluster":{
      "type":"string"
    },
    "kube_cluster":{
      "type":"string"
    },
    "volume":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	volumeDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "glusterfs_cluster":{
      "type":"string"
    },
    "kube_cluster":{
      "type":"string"
    },
    "volume":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *VolumeListRequest) InValid() (*gojsonschema.Result, error) {
	return volumeListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *VolumeCreateRequest) InValid() (*gojsonschema.Result, error) {
	return volumeCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *VolumeDeleteRequest) InValid() (*gojsonschema.Result, error) {
	return volumeDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
