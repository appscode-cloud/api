package ci

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var ypesVoidRequestSchema *gojsonschema.Schema
var jobCopyRequestSchema *gojsonschema.Schema
var jobDescribeRequestSchema *gojsonschema.Schema
var jobCreateRequestSchema *gojsonschema.Schema
var jobBuildRequestSchema *gojsonschema.Schema
var jobDeleteRequestSchema *gojsonschema.Schema

func init() {
	var err error
	ypesVoidRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	jobCopyRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	jobDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	jobCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	jobDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
}

func (m *ypesVoidRequest) InValid() (*gojsonschema.Result, error) {
	return ypesVoidRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *JobCopyRequest) InValid() (*gojsonschema.Result, error) {
	return jobCopyRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *JobDescribeRequest) InValid() (*gojsonschema.Result, error) {
	return jobDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *JobCreateRequest) InValid() (*gojsonschema.Result, error) {
	return jobCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *JobBuildRequest) InValid() (*gojsonschema.Result, error) {
	return jobBuildRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *JobDeleteRequest) InValid() (*gojsonschema.Result, error) {
	return jobDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
