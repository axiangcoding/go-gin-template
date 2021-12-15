// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "axiangcoding",
            "email": "axiangcoding@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/demo/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "demo，测试get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some params named param1",
                        "name": "param1",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "some params named param2",
                        "name": "param2",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/demo/post": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "demo，测试post",
                "parameters": [
                    {
                        "description": "some params json",
                        "name": "params",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/v1.Params"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/test/test-log": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "demo，测试post",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/login": {
            "post": {
                "tags": [
                    "user"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "register form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ErrJson"
                        }
                    }
                }
            }
        },
        "/api/v1/user/register": {
            "post": {
                "tags": [
                    "user"
                ],
                "summary": "User register",
                "parameters": [
                    {
                        "description": "register form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.RegisterForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.ErrJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ErrJson"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.ApiJson": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "app.ErrJson": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "string"
                }
            }
        },
        "v1.LoginForm": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "v1.Params": {
            "type": "object",
            "properties": {
                "param1": {
                    "type": "string"
                },
                "param2": {
                    "type": "string"
                }
            }
        },
        "v1.RegisterForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "token",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Gin Template API",
	Description: "An example of gin",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
