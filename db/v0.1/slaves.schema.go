package db

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var slaveAddRequestSchema *gojsonschema.Schema

func init() {
	var err error
	slaveAddRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "config": {
      "type": "string"
    },
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *SlaveAddRequest) IsValid() (*gojsonschema.Result, error) {
	return slaveAddRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SlaveAddRequest) IsRequest() {}

