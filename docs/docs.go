// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/apps/{guid}/features": {
            "get": {
                "description": "This endpoint retrieves the list of features for the specified app.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App Features"
                ],
                "summary": "List app features",
                "parameters": [
                    {
                        "type": "string",
                        "description": "App Guid",
                        "name": "guid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cf oauth-token",
                        "name": "cf-Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app_features.AppFeatureList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.Error"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.Error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/config.Error"
                        }
                    }
                }
            }
        },
        "/apps/{guid}/features/{name}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App Features"
                ],
                "summary": "Get an app feature",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cf oauth-token",
                        "name": "cf-Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "App Guid",
                        "name": "guid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "App Feature Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app_features.AppFeature"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.Error"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/config.Error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/config.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app_features.AppFeature": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "enabled": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "app_features.AppFeatureList": {
            "type": "object",
            "properties": {
                "pagination": {
                    "type": "object",
                    "properties": {
                        "first": {
                            "type": "object",
                            "properties": {
                                "href": {
                                    "type": "string"
                                }
                            }
                        },
                        "last": {
                            "type": "object",
                            "properties": {
                                "href": {
                                    "type": "string"
                                }
                            }
                        },
                        "next": {
                            "type": "object"
                        },
                        "previous": {
                            "type": "object"
                        },
                        "total_pages": {
                            "type": "integer"
                        },
                        "total_results": {
                            "type": "integer"
                        }
                    }
                },
                "resources": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app_features.AppFeature"
                    }
                }
            }
        },
        "config.Error": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.Errors"
                    }
                }
            }
        },
        "config.Errors": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "detail": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
