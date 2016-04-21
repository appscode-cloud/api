package bucket

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var bucketListRequestSchema *gojsonschema.Schema

func init() {
	var err error
	bucketListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "cloud_credential":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *BucketListRequest) InValid() (*gojsonschema.Result, error) {
	return bucketListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
