package loadbalancer

// Auto-generated. DO NOT EDIT.
import (
    "github.com/appscode/api/dtypes"
    "github.com/xeipuuv/gojsonschema"
    "log"
)
var listRequestSchema *gojsonschema.Schema
var deleteRequestSchema *gojsonschema.Schema
var createRequestSchema *gojsonschema.Schema
var describeRequestSchema *gojsonschema.Schema
var updateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	listRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	deleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "kind": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	createRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "LoadBalancerOptionsEntry": {
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "loadbalancerHTTPLoadBalancerRule": {
      "properties": {
        "SSL_secret_name": {
          "type": "string"
        },
        "backend": {
          "$ref": "#/definitions/loadbalancerLoadBalancerBackend"
        },
        "header_rule": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "path": {
          "type": "string"
        },
        "rewrite_rule": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "type": "object"
    },
    "loadbalancerLoadBalancer": {
      "properties": {
        "creation_timestamp": {
          "type": "string"
        },
        "kind": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "options": {
          "items": {
            "$ref": "#/definitions/LoadBalancerOptionsEntry"
          },
          "type": "array"
        },
        "spec": {
          "$ref": "#/definitions/loadbalancerSpec"
        },
        "status": {
          "$ref": "#/definitions/loadbalancerStatus"
        }
      },
      "type": "object"
    },
    "loadbalancerLoadBalancerBackend": {
      "properties": {
        "service_name": {
          "type": "string"
        },
        "service_port": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "loadbalancerLoadBalancerRule": {
      "properties": {
        "host": {
          "type": "string"
        },
        "http": {
          "items": {
            "$ref": "#/definitions/loadbalancerHTTPLoadBalancerRule"
          },
          "type": "array"
        },
        "tcp": {
          "items": {
            "$ref": "#/definitions/loadbalancerTCPLoadBalancerRule"
          },
          "type": "array"
        }
      },
      "type": "object"
    },
    "loadbalancerLoadBalancerStatus": {
      "properties": {
        "IP": {
          "type": "string"
        },
        "host": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "loadbalancerSpec": {
      "properties": {
        "backend": {
          "$ref": "#/definitions/loadbalancerHTTPLoadBalancerRule"
        },
        "rules": {
          "items": {
            "$ref": "#/definitions/loadbalancerLoadBalancerRule"
          },
          "type": "array"
        }
      },
      "type": "object"
    },
    "loadbalancerStatus": {
      "properties": {
        "status": {
          "items": {
            "$ref": "#/definitions/loadbalancerLoadBalancerStatus"
          },
          "type": "array"
        }
      },
      "type": "object"
    },
    "loadbalancerTCPLoadBalancerRule": {
      "properties": {
        "PEM_name": {
          "type": "string"
        },
        "backend": {
          "$ref": "#/definitions/loadbalancerLoadBalancerBackend"
        },
        "port": {
          "type": "string"
        },
        "secret_name": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "cluster": {
      "type": "string"
    },
    "load_balancer": {
      "$ref": "#/definitions/loadbalancerLoadBalancer"
    },
    "name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	describeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "kind": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	updateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "LoadBalancerOptionsEntry": {
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "loadbalancerHTTPLoadBalancerRule": {
      "properties": {
        "SSL_secret_name": {
          "type": "string"
        },
        "backend": {
          "$ref": "#/definitions/loadbalancerLoadBalancerBackend"
        },
        "header_rule": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "path": {
          "type": "string"
        },
        "rewrite_rule": {
          "items": {
            "type": "string"
          },
          "type": "array"
        }
      },
      "type": "object"
    },
    "loadbalancerLoadBalancer": {
      "properties": {
        "creation_timestamp": {
          "type": "string"
        },
        "kind": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "options": {
          "items": {
            "$ref": "#/definitions/LoadBalancerOptionsEntry"
          },
          "type": "array"
        },
        "spec": {
          "$ref": "#/definitions/loadbalancerSpec"
        },
        "status": {
          "$ref": "#/definitions/loadbalancerStatus"
        }
      },
      "type": "object"
    },
    "loadbalancerLoadBalancerBackend": {
      "properties": {
        "service_name": {
          "type": "string"
        },
        "service_port": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "loadbalancerLoadBalancerRule": {
      "properties": {
        "host": {
          "type": "string"
        },
        "http": {
          "items": {
            "$ref": "#/definitions/loadbalancerHTTPLoadBalancerRule"
          },
          "type": "array"
        },
        "tcp": {
          "items": {
            "$ref": "#/definitions/loadbalancerTCPLoadBalancerRule"
          },
          "type": "array"
        }
      },
      "type": "object"
    },
    "loadbalancerLoadBalancerStatus": {
      "properties": {
        "IP": {
          "type": "string"
        },
        "host": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "loadbalancerSpec": {
      "properties": {
        "backend": {
          "$ref": "#/definitions/loadbalancerHTTPLoadBalancerRule"
        },
        "rules": {
          "items": {
            "$ref": "#/definitions/loadbalancerLoadBalancerRule"
          },
          "type": "array"
        }
      },
      "type": "object"
    },
    "loadbalancerStatus": {
      "properties": {
        "status": {
          "items": {
            "$ref": "#/definitions/loadbalancerLoadBalancerStatus"
          },
          "type": "array"
        }
      },
      "type": "object"
    },
    "loadbalancerTCPLoadBalancerRule": {
      "properties": {
        "PEM_name": {
          "type": "string"
        },
        "backend": {
          "$ref": "#/definitions/loadbalancerLoadBalancerBackend"
        },
        "port": {
          "type": "string"
        },
        "secret_name": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "cluster": {
      "type": "string"
    },
    "load_balancer": {
      "$ref": "#/definitions/loadbalancerLoadBalancer"
    },
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ListRequest) IsValid() (*gojsonschema.Result, error) {
	return listRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ListRequest) IsRequest() {}

func (m *DeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return deleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DeleteRequest) IsRequest() {}

func (m *CreateRequest) IsValid() (*gojsonschema.Result, error) {
	return createRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CreateRequest) IsRequest() {}

func (m *DescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return describeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DescribeRequest) IsRequest() {}

func (m *UpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return updateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *UpdateRequest) IsRequest() {}

func (m *ListResponse) SetStatus(s *dtypes.Status) {
   m.Status = s
}

func (m *DescribeResponse) SetStatus(s *dtypes.Status) {
   m.Status = s
}

