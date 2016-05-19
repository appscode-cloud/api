package kubernetes

// Auto-generated. DO NOT EDIT.


import (
    "github.com/xeipuuv/gojsonschema"
    "log"
)
var eventRequestSchema *gojsonschema.Schema

func init() {
    	var err error
    	eventRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "EventRequestObjectMeta": {
      "properties": {
        "instance_id": {
          "type": "string"
        },
        "kind": {
          "type": "string"
        },
        "labels": {
          "items": {
            "$ref": "#/definitions/ObjectMetaLabelsEntry"
          },
          "type": "array"
        },
        "namespace": {
          "type": "string"
        },
        "pod_ip": {
          "type": "string"
        },
        "replication_controller": {
          "type": "string"
        },
        "service": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "type": "object"
    },
    "ObjectMetaLabelsEntry": {
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "cluster_name": {
      "type": "string"
    },
    "event_type": {
      "type": "string"
    },
    "kube_namespace": {
      "type": "string"
    },
    "kube_object_name": {
      "type": "string"
    },
    "kube_object_type": {
      "type": "string"
    },
    "metadata": {
      "$ref": "#/definitions/EventRequestObjectMeta"
    }
  },
  "type": "object"
}`))
	if err != nil {
    		log.Fatal(err)
    	}
    }

func (m *EventRequest) IsValid() (*gojsonschema.Result, error) {
	return eventRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *EventRequest) IsRequest() {}

func (m *EventResponse) SetStatus(s *dtypes.Status) {
   m.Status = s
}

