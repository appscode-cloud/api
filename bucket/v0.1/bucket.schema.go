package bucket

// Auto-generated. DO NOT EDIT.
import (
    "github.com/appscode/api/dtypes"
    "github.com/xeipuuv/gojsonschema"
    "log"
)
var bucketListRequestSchema *gojsonschema.Schema

func init() {
	var err error
	bucketListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cloud_credential": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *BucketListRequest) IsValid() (*gojsonschema.Result, error) {
	return bucketListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *BucketListRequest) IsRequest() {}

func (m *BucketListResponse) SetStatus(s *dtypes.Status) {
   m.Status = s
}

