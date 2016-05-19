package ci

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var masterDeleteRequestSchema *gojsonschema.Schema
var masterCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	masterDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	masterCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "type": "string"
    },
    "volume_id": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *MasterDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return masterDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *MasterDeleteRequest) IsRequest() {}

func (m *MasterCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return masterCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *MasterCreateRequest) IsRequest() {}

func (m *MasterDeleteResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

func (m *MasterCreateResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

