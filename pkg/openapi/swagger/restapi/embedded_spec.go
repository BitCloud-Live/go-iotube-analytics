// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "polydefi v1 api",
    "title": "polydefi",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@swagger.io"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "polydefi.bitcloud.live",
  "basePath": "/v1",
  "paths": {
    "/chart/{days}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "data"
        ],
        "summary": "Get chart defi data",
        "operationId": "getChartData",
        "parameters": [
          {
            "type": "integer",
            "description": "Number of days",
            "name": "days",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ChartData"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/data": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "data"
        ],
        "summary": "Get all defi data",
        "operationId": "getAllData",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/AllData"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AllData": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/DefiData"
      }
    },
    "ApiResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "x-omitempty": false
        },
        "status": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "ChartData": {
      "type": "object"
    },
    "DefiData": {
      "type": "object",
      "properties": {
        "category": {
          "description": "Category",
          "type": "string",
          "x-omitempty": false
        },
        "chain": {
          "description": "Name",
          "type": "string",
          "x-omitempty": false
        },
        "contractNum": {
          "description": "Contract Num",
          "type": "integer",
          "x-omitempty": false
        },
        "holders": {
          "description": "Holders",
          "type": "integer",
          "x-omitempty": false
        },
        "holdersChange24hNum": {
          "description": "Holders Change 24h",
          "type": "integer",
          "x-omitempty": false
        },
        "lastUpdated": {
          "description": "Last Updated",
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "lockedUsd": {
          "description": "Locked Usd",
          "type": "integer",
          "x-omitempty": false
        },
        "marketCap": {
          "description": "Market Cap",
          "type": "number",
          "x-omitempty": false
        },
        "marketCapChange24h": {
          "description": "Market Cap Change 24h",
          "type": "number",
          "x-omitempty": false
        },
        "name": {
          "description": "Name",
          "type": "string",
          "x-omitempty": false
        },
        "price": {
          "description": "Price",
          "type": "number",
          "x-omitempty": false
        },
        "priceChange24h": {
          "description": "Price Percent Change 24h",
          "type": "number",
          "x-omitempty": false
        },
        "token": {
          "description": "Token",
          "type": "string"
        },
        "tvlPercentChange24h": {
          "description": "TVL Percent Change 24h",
          "type": "number",
          "x-omitempty": false
        },
        "verified": {
          "description": "Verified",
          "type": "integer",
          "x-omitempty": false
        },
        "volume": {
          "description": "Volume",
          "type": "integer",
          "x-omitempty": false
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "polydefi v1 api",
    "title": "polydefi",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@swagger.io"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "polydefi.bitcloud.live",
  "basePath": "/v1",
  "paths": {
    "/chart/{days}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "data"
        ],
        "summary": "Get chart defi data",
        "operationId": "getChartData",
        "parameters": [
          {
            "type": "integer",
            "description": "Number of days",
            "name": "days",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ChartData"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/data": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "data"
        ],
        "summary": "Get all defi data",
        "operationId": "getAllData",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/AllData"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AllData": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/DefiData"
      }
    },
    "ApiResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "x-omitempty": false
        },
        "status": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "ChartData": {
      "type": "object"
    },
    "DefiData": {
      "type": "object",
      "properties": {
        "category": {
          "description": "Category",
          "type": "string",
          "x-omitempty": false
        },
        "chain": {
          "description": "Name",
          "type": "string",
          "x-omitempty": false
        },
        "contractNum": {
          "description": "Contract Num",
          "type": "integer",
          "x-omitempty": false
        },
        "holders": {
          "description": "Holders",
          "type": "integer",
          "x-omitempty": false
        },
        "holdersChange24hNum": {
          "description": "Holders Change 24h",
          "type": "integer",
          "x-omitempty": false
        },
        "lastUpdated": {
          "description": "Last Updated",
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "lockedUsd": {
          "description": "Locked Usd",
          "type": "integer",
          "x-omitempty": false
        },
        "marketCap": {
          "description": "Market Cap",
          "type": "number",
          "x-omitempty": false
        },
        "marketCapChange24h": {
          "description": "Market Cap Change 24h",
          "type": "number",
          "x-omitempty": false
        },
        "name": {
          "description": "Name",
          "type": "string",
          "x-omitempty": false
        },
        "price": {
          "description": "Price",
          "type": "number",
          "x-omitempty": false
        },
        "priceChange24h": {
          "description": "Price Percent Change 24h",
          "type": "number",
          "x-omitempty": false
        },
        "token": {
          "description": "Token",
          "type": "string"
        },
        "tvlPercentChange24h": {
          "description": "TVL Percent Change 24h",
          "type": "number",
          "x-omitempty": false
        },
        "verified": {
          "description": "Verified",
          "type": "integer",
          "x-omitempty": false
        },
        "volume": {
          "description": "Volume",
          "type": "integer",
          "x-omitempty": false
        }
      }
    }
  }
}`))
}
