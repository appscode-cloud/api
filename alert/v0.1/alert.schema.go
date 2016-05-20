package alert

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var deleteRequestSchema *gojsonschema.Schema
var listRequestSchema *gojsonschema.Schema
var acknowledgeRequestSchema *gojsonschema.Schema
var updateRequestSchema *gojsonschema.Schema
var createRequestSchema *gojsonschema.Schema
var notificationRequestSchema *gojsonschema.Schema

func init() {
	var err error
	deleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "alertSpec": {
      "properties": {
        "cluster": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "object_name": {
          "type": "string"
        },
        "object_type": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "phid": {
      "type": "string"
    },
    "spec": {
      "$ref": "#/definitions/alertSpec"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	listRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "alertSpec": {
      "properties": {
        "cluster": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "object_name": {
          "type": "string"
        },
        "object_type": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "spec": {
      "$ref": "#/definitions/alertSpec"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	acknowledgeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "comment": {
      "type": "string"
    },
    "phid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	updateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "alertAlertSpec": {
      "properties": {
        "alert_interval": {
          "type": "integer"
        },
        "critical_method": {
          "type": "integer"
        },
        "critical_user": {
          "type": "string"
        },
        "warning_method": {
          "type": "integer"
        },
        "warning_user": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "alertIcingaService": {
      "properties": {
        "check_command": {
          "type": "string"
        },
        "check_interval": {
          "type": "integer"
        },
        "critical_condition": {
          "type": "string"
        },
        "formula_r": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "query": {
          "additionalProperties": {
            "type": "string"
          },
          "type": "object"
        },
        "warning_condition": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "alertSpec": {
      "properties": {
        "cluster": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "object_name": {
          "type": "string"
        },
        "object_type": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "alert_spec": {
      "$ref": "#/definitions/alertAlertSpec"
    },
    "icinga_service": {
      "$ref": "#/definitions/alertIcingaService"
    },
    "icinga_service_phid": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "phid": {
      "type": "string"
    },
    "spec": {
      "$ref": "#/definitions/alertSpec"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	createRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "alertAlertSpec": {
      "properties": {
        "alert_interval": {
          "type": "integer"
        },
        "critical_method": {
          "type": "integer"
        },
        "critical_user": {
          "type": "string"
        },
        "warning_method": {
          "type": "integer"
        },
        "warning_user": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "alertIcingaService": {
      "properties": {
        "check_command": {
          "type": "string"
        },
        "check_interval": {
          "type": "integer"
        },
        "critical_condition": {
          "type": "string"
        },
        "formula_r": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "query": {
          "additionalProperties": {
            "type": "string"
          },
          "type": "object"
        },
        "warning_condition": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "alertSpec": {
      "properties": {
        "cluster": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "object_name": {
          "type": "string"
        },
        "object_type": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "alert_spec": {
      "$ref": "#/definitions/alertAlertSpec"
    },
    "icinga_service": {
      "$ref": "#/definitions/alertIcingaService"
    },
    "icinga_service_phid": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "spec": {
      "$ref": "#/definitions/alertSpec"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	notificationRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "alert_phid": {
      "type": "string"
    },
    "host_name": {
      "type": "string"
    },
    "service_output": {
      "type": "string"
    },
    "state": {
      "type": "string"
    },
    "state_type": {
      "type": "string"
    },
    "time": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *DeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return deleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DeleteRequest) IsRequest() {}

func (m *ListRequest) IsValid() (*gojsonschema.Result, error) {
	return listRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ListRequest) IsRequest() {}

func (m *AcknowledgeRequest) IsValid() (*gojsonschema.Result, error) {
	return acknowledgeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AcknowledgeRequest) IsRequest() {}

func (m *UpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return updateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *UpdateRequest) IsRequest() {}

func (m *CreateRequest) IsValid() (*gojsonschema.Result, error) {
	return createRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CreateRequest) IsRequest() {}

func (m *NotificationRequest) IsValid() (*gojsonschema.Result, error) {
	return notificationRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *NotificationRequest) IsRequest() {}

func (m *ListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
