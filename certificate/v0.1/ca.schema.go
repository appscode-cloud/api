package certificate

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var cACreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	cACreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "csr": {
      "type": "string"
    },
    "name": {
      "type": "string"
    }
  },
  "title": "Use specific requests for protos",
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *CACreateRequest) IsValid() (*gojsonschema.Result, error) {
	return cACreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CACreateRequest) IsRequest() {}

func (m *CACreateResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}

