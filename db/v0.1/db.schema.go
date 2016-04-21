package db

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var deleteRequestSchema *gojsonschema.Schema
var describeRequestSchema *gojsonschema.Schema
var ypesVoidRequestSchema *gojsonschema.Schema
var backupRequestSchema *gojsonschema.Schema
var createRequestSchema *gojsonschema.Schema
var snapshotListRequestSchema *gojsonschema.Schema
var restoreRequestSchema *gojsonschema.Schema

func init() {
	var err error
	deleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cluster":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "type":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	describeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cluster":{
      "type":"string"
    },
    "name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	ypesVoidRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	backupRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "auth_secret_name":{
      "type":"string"
    },
    "bucket_name":{
      "type":"string"
    },
    "cluster":{
      "type":"string"
    },
    "credential":{
      "type":"string"
    },
    "force":{
      "type":"boolean"
    },
    "name":{
      "type":"string"
    },
    "region":{
      "type":"string"
    },
    "secret_name":{
      "type":"string"
    },
    "snapshot_name":{
      "type":"string"
    },
    "type":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	createRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "bucket_name":{
      "type":"string"
    },
    "cluster":{
      "type":"string"
    },
    "credential":{
      "type":"string"
    },
    "gfs_disk":{
      "type":"string"
    },
    "gfs_endpoint":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "node":{
      "type":"integer"
    },
    "pv":{
      "type":"string"
    },
    "pv_size":{
      "type":"integer"
    },
    "region":{
      "type":"string"
    },
    "sku":{
      "type":"string"
    },
    "type":{
      "type":"string"
    },
    "version":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	snapshotListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	restoreRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "bucket_name":{
      "type":"string"
    },
    "cluster":{
      "type":"string"
    },
    "credential":{
      "type":"string"
    },
    "force":{
      "type":"boolean"
    },
    "name":{
      "type":"string"
    },
    "node":{
      "type":"integer"
    },
    "pv":{
      "type":"string"
    },
    "pv_size":{
      "type":"integer"
    },
    "region":{
      "type":"string"
    },
    "secret_name":{
      "type":"string"
    },
    "sku":{
      "type":"string"
    },
    "snapshot_phid":{
      "type":"string"
    },
    "type":{
      "type":"string"
    },
    "version":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *DeleteRequest) InValid() (*gojsonschema.Result, error) {
	return deleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DescribeRequest) InValid() (*gojsonschema.Result, error) {
	return describeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ypesVoidRequest) InValid() (*gojsonschema.Result, error) {
	return ypesVoidRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *BackupRequest) InValid() (*gojsonschema.Result, error) {
	return backupRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CreateRequest) InValid() (*gojsonschema.Result, error) {
	return createRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SnapshotListRequest) InValid() (*gojsonschema.Result, error) {
	return snapshotListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *RestoreRequest) InValid() (*gojsonschema.Result, error) {
	return restoreRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
