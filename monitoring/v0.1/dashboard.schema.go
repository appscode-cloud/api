package monitoring

import (
	"github.com/xeipuuv/gojsonschema"
	// "log"
)

// Auto-generated. DO NOT EDIT.

var dashboardCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	dashboardCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "cluster_name":{
      "type":"string"
    },
    "namespace":{
      "type":"string"
    },
    "object_name":{
      "type":"string"
    },
    "type":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		// log.Fatal(err)
	}
}

func (m *DashboardCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return dashboardCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
