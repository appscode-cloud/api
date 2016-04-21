package kubernetes

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var clusterInstanceListRequestSchema *gojsonschema.Schema
var clusterScaleRequestSchema *gojsonschema.Schema
var clusterUpdateRequestSchema *gojsonschema.Schema
var clusterDeleteRequestSchema *gojsonschema.Schema
var clusterDescribeRequestSchema *gojsonschema.Schema
var clusterStartupScriptRequestSchema *gojsonschema.Schema
var clusterCreateRequestSchema *gojsonschema.Schema
var clusterClientConfigRequestSchema *gojsonschema.Schema

func init() {
	var err error
	clusterInstanceListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cluster_name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	clusterScaleRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "name":{
      "type":"string"
    },
    "node_changes":{
      "items":{
        "$ref":"#/definitions/ClusterScaleRequestNodeChangesEntry"
      },
      "type":"array"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	clusterUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "kube_starter_version":{
      "type":"string"
    },
    "kube_version":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "saltbase_version":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	clusterDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "name":{
      "type":"string"
    },
    "release_reserved_ip":{
      "type":"boolean"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	clusterDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	clusterStartupScriptRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "role":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	clusterCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cloud_credential":{
      "type":"string"
    },
    "cloud_credential_data":{
      "items":{
        "$ref":"#/definitions/ClusterCreateRequestCloudCredentialDataEntry"
      },
      "type":"array"
    },
    "kube_starter_version":{
      "type":"string"
    },
    "kube_version":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "node_set":{
      "items":{
        "$ref":"#/definitions/ClusterCreateRequestNodeSetEntry"
      },
      "type":"array"
    },
    "provider":{
      "type":"string"
    },
    "saltbase_version":{
      "type":"string"
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
	clusterClientConfigRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ClusterInstanceListRequest) InValid() (*gojsonschema.Result, error) {
	return clusterInstanceListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterScaleRequest) InValid() (*gojsonschema.Result, error) {
	return clusterScaleRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterUpdateRequest) InValid() (*gojsonschema.Result, error) {
	return clusterUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterDeleteRequest) InValid() (*gojsonschema.Result, error) {
	return clusterDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterDescribeRequest) InValid() (*gojsonschema.Result, error) {
	return clusterDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterStartupScriptRequest) InValid() (*gojsonschema.Result, error) {
	return clusterStartupScriptRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterCreateRequest) InValid() (*gojsonschema.Result, error) {
	return clusterCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterClientConfigRequest) InValid() (*gojsonschema.Result, error) {
	return clusterClientConfigRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
