package health

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var voidRequestSchema *gojsonschema.Schema

func init() {
	var err error
	voidRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *VoidRequest) InValid() (*gojsonschema.Result, error) {
	return voidRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
