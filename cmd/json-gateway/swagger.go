package main

var swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "storage.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v1/all": {
      "post": {
        "summary": "All returns all the namespaces for a given environment.",
        "operationId": "All",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/AllReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AllRequest"
            }
          }
        ],
        "tags": [
          "ElwinStorage"
        ]
      }
    },
    "/api/v1/create": {
      "post": {
        "summary": "Create creates a namespace in the given environment.",
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/CreateReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateRequest"
            }
          }
        ],
        "tags": [
          "ElwinStorage"
        ]
      }
    },
    "/api/v1/delete": {
      "post": {
        "summary": "Delete deletes the namespace from the given environment.",
        "operationId": "Delete",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/DeleteReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/DeleteRequest"
            }
          }
        ],
        "tags": [
          "ElwinStorage"
        ]
      }
    },
    "/api/v1/experiment-intake": {
      "post": {
        "summary": "ExperimentIntake takes a request from a web form and creates the\nexperiment in the data store.",
        "operationId": "ExperimentIntake",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/ExperimentIntakeReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ExperimentIntakeRequest"
            }
          }
        ],
        "tags": [
          "ElwinStorage"
        ]
      }
    },
    "/api/v1/read": {
      "post": {
        "summary": "Read returns the namespace matching the supplied name from the given\nenvironment.",
        "operationId": "Read",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/ReadReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ReadRequest"
            }
          }
        ],
        "tags": [
          "ElwinStorage"
        ]
      }
    },
    "/api/v1/update": {
      "post": {
        "summary": "Update replaces the namespace in the given environment with the namespace\nsupplied.",
        "operationId": "Update",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/UpdateReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UpdateRequest"
            }
          }
        ],
        "tags": [
          "ElwinStorage"
        ]
      }
    }
  },
  "definitions": {
    "AllReply": {
      "type": "object",
      "properties": {
        "namespaces": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Namespace"
          }
        }
      },
      "title": "The response message containing the Namespaces"
    },
    "AllRequest": {
      "type": "object",
      "properties": {
        "environment": {
          "$ref": "#/definitions/Environment"
        }
      },
      "description": "AllRequest returns all the experiments from the given environment."
    },
    "CreateReply": {
      "type": "object",
      "properties": {
        "namespace": {
          "$ref": "#/definitions/Namespace"
        }
      },
      "description": "CreateReply response containing the newly created Namespace."
    },
    "CreateRequest": {
      "type": "object",
      "properties": {
        "namespace": {
          "$ref": "#/definitions/Namespace"
        },
        "environment": {
          "$ref": "#/definitions/Environment"
        }
      },
      "description": "CreateRequest request message to create a new namespace in an environment."
    },
    "DeleteReply": {
      "type": "object",
      "properties": {
        "namespace": {
          "$ref": "#/definitions/Namespace"
        }
      },
      "description": "DeleteReply response containing the deleted namespace."
    },
    "DeleteRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "environment": {
          "$ref": "#/definitions/Environment"
        }
      },
      "description": "DeleteRequest request message to delete an existing namespace from an\nenvironment."
    },
    "Environment": {
      "type": "string",
      "enum": [
        "Staging",
        "Production"
      ],
      "default": "Staging",
      "title": "Environment structure"
    },
    "Experiment": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "segments": {
          "type": "string",
          "format": "byte"
        },
        "params": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Param"
          }
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "detailName": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "Experiment structure"
    },
    "ExperimentIntakeReply": {
      "type": "object"
    },
    "ExperimentIntakeRequest": {
      "type": "object",
      "properties": {
        "metadata": {
          "$ref": "#/definitions/ExperimentMetadata"
        },
        "namespace": {
          "$ref": "#/definitions/Namespace"
        }
      },
      "title": "ExperimentIntakeRequest creates an experiment in the database and sends a notification for reviewers"
    },
    "ExperimentMetadata": {
      "type": "object",
      "properties": {
        "userID": {
          "type": "string",
          "format": "string"
        },
        "programManagerID": {
          "type": "string",
          "format": "string"
        },
        "productManagerID": {
          "type": "string",
          "format": "string"
        },
        "hypothesis": {
          "type": "string",
          "format": "string"
        },
        "kpi": {
          "type": "string",
          "format": "string"
        },
        "timeBound": {
          "type": "boolean",
          "format": "boolean"
        },
        "plannedStartTime": {
          "type": "string",
          "format": "string"
        },
        "plannedEndTime": {
          "type": "string",
          "format": "string"
        },
        "actualStartTime": {
          "type": "string",
          "format": "string"
        },
        "actualEndTime": {
          "type": "string",
          "format": "string"
        },
        "actionPlanNegative": {
          "type": "string",
          "format": "string"
        },
        "actionPlanNeutral": {
          "type": "string",
          "format": "string"
        },
        "experimentType": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "ExperimentMetadata all the junk that elwin doesn't care about"
    },
    "Namespace": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "experiments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Experiment"
          }
        }
      },
      "title": "Namespace structure"
    },
    "Param": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "value": {
          "$ref": "#/definitions/Value"
        },
        "id": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "Param structure"
    },
    "ReadReply": {
      "type": "object",
      "properties": {
        "namespace": {
          "$ref": "#/definitions/Namespace"
        }
      },
      "description": "ReadReply response containing the namespace requested."
    },
    "ReadRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "environment": {
          "$ref": "#/definitions/Environment"
        }
      },
      "description": "ReadRequest request message to get a namespace by name."
    },
    "UpdateReply": {
      "type": "object",
      "properties": {
        "namespace": {
          "$ref": "#/definitions/Namespace"
        }
      },
      "description": "UpdateReply response containing the updated namespace."
    },
    "UpdateRequest": {
      "type": "object",
      "properties": {
        "namespace": {
          "$ref": "#/definitions/Namespace"
        },
        "environment": {
          "$ref": "#/definitions/Environment"
        }
      },
      "description": "UpdateRequest request message to update an existing namespace in an\nenvironment."
    },
    "Value": {
      "type": "object",
      "properties": {
        "choices": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "weights": {
          "type": "array",
          "items": {
            "type": "number",
            "format": "double"
          }
        }
      },
      "title": "Value structure"
    }
  }
}`
