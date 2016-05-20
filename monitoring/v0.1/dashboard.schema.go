package monitoring

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var dashboardCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	dashboardCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    },
    "object_name": {
      "type": "string"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *DashboardCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return dashboardCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DashboardCreateRequest) IsRequest() {}

func (m *DashboardCreateResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
