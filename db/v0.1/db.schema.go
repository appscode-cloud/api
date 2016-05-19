package db

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var deleteRequestSchema *gojsonschema.Schema
var describeRequestSchema *gojsonschema.Schema
var backupRequestSchema *gojsonschema.Schema
var createRequestSchema *gojsonschema.Schema
var snapshotListRequestSchema *gojsonschema.Schema
var restoreRequestSchema *gojsonschema.Schema

func init() {
	var err error
	deleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	describeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
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
	backupRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "auth_secret_name": {
      "type": "string"
    },
    "bucket_name": {
      "type": "string"
    },
    "cluster": {
      "type": "string"
    },
    "credential": {
      "type": "string"
    },
    "force": {
      "type": "boolean"
    },
    "name": {
      "type": "string"
    },
    "region": {
      "type": "string"
    },
    "secret_name": {
      "type": "string"
    },
    "snapshot_name": {
      "type": "string"
    },
    "type": {
      "type": "string"
    },
    "wal": {
      "type": "boolean"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	createRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "auth_secret_name": {
      "type": "string"
    },
    "bucket_name": {
      "type": "string"
    },
    "cluster": {
      "type": "string"
    },
    "credential": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "node": {
      "type": "integer"
    },
    "pv": {
      "type": "string"
    },
    "pv_size": {
      "type": "integer"
    },
    "region": {
      "type": "string"
    },
    "sku": {
      "type": "string"
    },
    "type": {
      "type": "string"
    },
    "version": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	snapshotListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	restoreRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "auth_secret_name": {
      "type": "string"
    },
    "bucket_name": {
      "type": "string"
    },
    "cluster": {
      "type": "string"
    },
    "credential": {
      "type": "string"
    },
    "force": {
      "type": "boolean"
    },
    "name": {
      "type": "string"
    },
    "node": {
      "type": "integer"
    },
    "pv": {
      "type": "string"
    },
    "pv_size": {
      "type": "integer"
    },
    "region": {
      "type": "string"
    },
    "secret_name": {
      "type": "string"
    },
    "sku": {
      "type": "string"
    },
    "snapshot_phid": {
      "type": "string"
    },
    "type": {
      "type": "string"
    },
    "version": {
      "type": "string"
    },
    "wal": {
      "type": "boolean"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *DeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return deleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DeleteRequest) IsRequest() {}

func (m *DescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return describeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DescribeRequest) IsRequest() {}

func (m *BackupRequest) IsValid() (*gojsonschema.Result, error) {
	return backupRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *BackupRequest) IsRequest() {}

func (m *CreateRequest) IsValid() (*gojsonschema.Result, error) {
	return createRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CreateRequest) IsRequest() {}

func (m *SnapshotListRequest) IsValid() (*gojsonschema.Result, error) {
	return snapshotListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SnapshotListRequest) IsRequest() {}

func (m *RestoreRequest) IsValid() (*gojsonschema.Result, error) {
	return restoreRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *RestoreRequest) IsRequest() {}

func (m *ListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

func (m *SnapshotListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

func (m *DescribeResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

