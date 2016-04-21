package monitoring

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var dashboardCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	dashboardCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
		log.Fatal(err)
	}
}

func (m *DashboardCreateRequest) InValid() (*gojsonschema.Result, error) {
	return dashboardCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
