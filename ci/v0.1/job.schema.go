package ci

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var jobDescribeRequestSchema *gojsonschema.Schema
var jobDeleteRequestSchema *gojsonschema.Schema
var jobCopyRequestSchema *gojsonschema.Schema
var jobCreateRequestSchema *gojsonschema.Schema
var jobBuildRequestSchema *gojsonschema.Schema

func init() {
	var err error
	jobDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	jobDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "name":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	jobCopyRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "destination":{
      "type":"string"
    },
    "source":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	jobCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "name":{
      "type":"string"
    },
    "sh_file":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	jobBuildRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema":"http://json-schema.org/draft-04/schema#",
  "properties":{
    "name":{
      "type":"string"
    },
    "param":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *JobDescribeRequest) InValid() (*gojsonschema.Result, error) {
	return jobDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *JobDeleteRequest) InValid() (*gojsonschema.Result, error) {
	return jobDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *JobCopyRequest) InValid() (*gojsonschema.Result, error) {
	return jobCopyRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *JobCreateRequest) InValid() (*gojsonschema.Result, error) {
	return jobCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *JobBuildRequest) InValid() (*gojsonschema.Result, error) {
	return jobBuildRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
