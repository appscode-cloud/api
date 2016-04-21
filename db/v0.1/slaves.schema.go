package db

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var slaveAddRequestSchema *gojsonschema.Schema

func init() {
	var err error
	slaveAddRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cluster":{
      "type":"string"
    },
    "config":{
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
}

func (m *SlaveAddRequest) InValid() (*gojsonschema.Result, error) {
	return slaveAddRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
