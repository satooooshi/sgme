// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/addServiceApidocUrl/{ns}/{svcname}/{apidocurl}": {
            "get": {
                "description": "addServiceApidocUrl",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "addServiceApidocUrl",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "service name",
                        "name": "svcname",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "api document url",
                        "name": "apidocurl",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    }
                }
            }
        },
        "/circuitBreaker/{ns}/{svcname}": {
            "get": {
                "description": "circuitBreaker defines circuit breaker policy that applies to traffic intended for a service after routing has occurred.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "circuitBreaker",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "service name",
                        "name": "svcname",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    }
                }
            }
        },
        "/discoverGateway/{ns}": {
            "get": {
                "description": "discoverGateway",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "discoverGateway",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    }
                }
            }
        },
        "/discoverGatewayPort/{ns}": {
            "get": {
                "description": "discoverGatewayPort",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "discoverGatewayPort",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    }
                }
            }
        },
        "/discoverService/{ns}/{svcname}": {
            "get": {
                "description": "discoverService",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "discoverService",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "service name",
                        "name": "svcname",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    }
                }
            }
        },
        "/discoverServiceApidocUrl/{ns}/{svcname}": {
            "get": {
                "description": "discoverServiceApidocUrl",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "discoverServiceApidocUrl",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "service name",
                        "name": "svcname",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    }
                }
            }
        },
        "/discoverServices{ns}": {
            "get": {
                "description": "discoverServices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "discoverServices",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    }
                }
            }
        },
        "/registerGateway/{ns}/{ip}": {
            "get": {
                "description": "registerGateway",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "registerGateway",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "gateway external ip address",
                        "name": "ip",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    }
                }
            }
        },
        "/registerService/{ns}/{svcname}": {
            "get": {
                "description": "registerService",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "registerService",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "service name",
                        "name": "svcname",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    }
                }
            }
        },
        "/test/": {
            "get": {
                "description": "テスト用APIの詳細",
                "consumes": [
                    "application/x-json-stream"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "必須ではありません。",
                        "name": "none",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "0.0.0.0:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "SGME API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}