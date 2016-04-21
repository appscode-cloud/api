package alert

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

var listRequestSchema *gojsonschema.Schema
var deleteRequestSchema *gojsonschema.Schema
var createRequestSchema *gojsonschema.Schema
var notificationRequestSchema *gojsonschema.Schema
var updateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	listRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "definitions":{
    "alertSpec":{
      "properties":{
        "cluster":{
          "type":"string"
        },
        "namespace":{
          "type":"string"
        },
        "object_name":{
          "type":"string"
        },
        "object_type":{
          "type":"string"
        }
      },
      "type":"object"
    }
  },
  "properties":{
    "spec":{
      "$ref":"#/definitions/alertSpec"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	deleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "definitions":{
    "alertSpec":{
      "properties":{
        "cluster":{
          "type":"string"
        },
        "namespace":{
          "type":"string"
        },
        "object_name":{
          "type":"string"
        },
        "object_type":{
          "type":"string"
        }
      },
      "type":"object"
    }
  },
  "properties":{
    "phid":{
      "type":"string"
    },
    "spec":{
      "$ref":"#/definitions/alertSpec"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	createRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "definitions":{
    "alertAlertSpec":{
      "properties":{
        "alert_interval":{
          "type":"integer"
        },
        "critical_condition":{
          "type":"string"
        },
        "critical_method":{
          "type":"integer"
        },
        "critical_user":{
          "type":"string"
        },
        "warning_condition":{
          "type":"string"
        },
        "warning_method":{
          "type":"integer"
        },
        "warning_user":{
          "type":"string"
        }
      },
      "type":"object"
    },
    "alertMatrix":{
      "properties":{
        "command":{
          "type":"string"
        },
        "formula":{
          "type":"string"
        },
        "name":{
          "type":"string"
        },
        "query":{
          "items":{
            "$ref":"#/definitions/MatrixQueryEntry"
          },
          "type":"array"
        }
      },
      "type":"object"
    },
    "alertSpec":{
      "properties":{
        "cluster":{
          "type":"string"
        },
        "namespace":{
          "type":"string"
        },
        "object_name":{
          "type":"string"
        },
        "object_type":{
          "type":"string"
        }
      },
      "type":"object"
    }
  },
  "properties":{
    "alert_spec":{
      "$ref":"#/definitions/alertAlertSpec"
    },
    "check_interval":{
      "type":"integer"
    },
    "matrix":{
      "$ref":"#/definitions/alertMatrix"
    },
    "matrix_phid":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "spec":{
      "$ref":"#/definitions/alertSpec"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	notificationRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "alert_phid":{
      "type":"string"
    },
    "host_name":{
      "type":"string"
    },
    "service_output":{
      "type":"string"
    },
    "state":{
      "type":"string"
    },
    "state_type":{
      "type":"string"
    },
    "time":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	updateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "definitions":{
    "alertAlertSpec":{
      "properties":{
        "alert_interval":{
          "type":"integer"
        },
        "critical_condition":{
          "type":"string"
        },
        "critical_method":{
          "type":"integer"
        },
        "critical_user":{
          "type":"string"
        },
        "warning_condition":{
          "type":"string"
        },
        "warning_method":{
          "type":"integer"
        },
        "warning_user":{
          "type":"string"
        }
      },
      "type":"object"
    },
    "alertMatrix":{
      "properties":{
        "command":{
          "type":"string"
        },
        "formula":{
          "type":"string"
        },
        "name":{
          "type":"string"
        },
        "query":{
          "items":{
            "$ref":"#/definitions/MatrixQueryEntry"
          },
          "type":"array"
        }
      },
      "type":"object"
    },
    "alertSpec":{
      "properties":{
        "cluster":{
          "type":"string"
        },
        "namespace":{
          "type":"string"
        },
        "object_name":{
          "type":"string"
        },
        "object_type":{
          "type":"string"
        }
      },
      "type":"object"
    }
  },
  "properties":{
    "alert_spec":{
      "$ref":"#/definitions/alertAlertSpec"
    },
    "check_interval":{
      "type":"integer"
    },
    "matrix":{
      "$ref":"#/definitions/alertMatrix"
    },
    "matrix_phid":{
      "type":"string"
    },
    "name":{
      "type":"string"
    },
    "phid":{
      "type":"string"
    },
    "spec":{
      "$ref":"#/definitions/alertSpec"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ListRequest) InValid() (*gojsonschema.Result, error) {
	return listRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DeleteRequest) InValid() (*gojsonschema.Result, error) {
	return deleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CreateRequest) InValid() (*gojsonschema.Result, error) {
	return createRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *NotificationRequest) InValid() (*gojsonschema.Result, error) {
	return notificationRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *UpdateRequest) InValid() (*gojsonschema.Result, error) {
	return updateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
