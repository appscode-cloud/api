package kubernetes

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var eventRequestSchema *gojsonschema.Schema

func init() {
	var err error
	eventRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "definitions":{
    "EventRequestObjectMeta":{
      "properties":{
        "instance_id":{
          "type":"string"
        },
        "kind":{
          "type":"string"
        },
        "labels":{
          "items":{
            "$ref":"#/definitions/ObjectMetaLabelsEntry"
          },
          "type":"array"
        },
        "namespace":{
          "type":"string"
        },
        "pod_ip":{
          "type":"string"
        },
        "replication_controller":{
          "type":"string"
        },
        "service":{
          "items":{
            "format":"string",
            "type":"string"
          },
          "type":"array"
        }
      },
      "type":"object"
    },
    "ObjectMetaLabelsEntry":{
      "properties":{
        "key":{
          "type":"string"
        },
        "value":{
          "type":"string"
        }
      },
      "type":"object"
    }
  },
  "properties":{
    "cluster_name":{
      "type":"string"
    },
    "event_type":{
      "type":"string"
    },
    "kube_namespace":{
      "type":"string"
    },
    "kube_object_name":{
      "type":"string"
    },
    "kube_object_type":{
      "type":"string"
    },
    "metadata":{
      "$ref":"#/definitions/EventRequestObjectMeta"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *EventRequest) InValid() (*gojsonschema.Result, error) {
	return eventRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
