package glusterfs

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var clusterDescribeRequestSchema *gojsonschema.Schema
var clusterListRequestSchema *gojsonschema.Schema
var clusterDeleteRequestSchema *gojsonschema.Schema
var clusterCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	clusterDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "kube_cluster": {
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
	clusterListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "kube_cluster": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	clusterDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "kube_cluster": {
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
	clusterCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "disks": {
      "items": {
        "type": "string"
      },
      "type": "array"
    },
    "enable_bacula": {
      "type": "integer"
    },
    "kube_cluster": {
      "type": "string"
    },
    "mood": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "node": {
      "type": "integer"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ClusterDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterDescribeRequest) IsRequest() {}

func (m *ClusterListRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterListRequest) IsRequest() {}

func (m *ClusterDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterDeleteRequest) IsRequest() {}

func (m *ClusterCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterCreateRequest) IsRequest() {}
