package ssh

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var secureShellGetRequestSchema *gojsonschema.Schema

func init() {
	var err error
	secureShellGetRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "type": "string"
    },
    "cluster_namespace": {
      "type": "string"
    },
    "instance_name": {
      "type": "string"
    },
    "jenkins_namespace": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *SecureShellGetRequest) IsValid() (*gojsonschema.Result, error) {
	return secureShellGetRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SecureShellGetRequest) IsRequest() {}

func (m *SecureShellGetResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

