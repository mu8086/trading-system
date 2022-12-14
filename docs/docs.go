// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2022-09-10 12:40:02.721096 +0800 CST m=+0.068880263

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
        "termsOfService": "https://github.com/mu8086/trading-system",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/orders": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "取得多個訂單",
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Order"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增訂單",
                "parameters": [
                    {
                        "maxLength": 255,
                        "minLength": 3,
                        "type": "string",
                        "description": "訂單擁有者",
                        "name": "owner",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "訂單種類, 1為買, 2為賣",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "交易數量",
                        "name": "quantity",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "交易價格",
                        "name": "price",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "價格原則, 1為限價, 2為市價",
                        "name": "price_policy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Order"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/orders/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "取得單個訂單",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "訂單編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Order"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "更新訂單",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "訂單編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 255,
                        "type": "string",
                        "description": "訂單擁有者",
                        "name": "owner",
                        "in": "query"
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "訂單種類, 1為買, 2為賣",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "交易數量",
                        "name": "quantity",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "交易價格",
                        "name": "price",
                        "in": "query"
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "價格原則, 1為限價, 2為市價",
                        "name": "price_policy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Order"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "刪除訂單",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "訂單編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Order"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errcode.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "details": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.Order": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "price_policy": {
                    "description": "limit price / market price",
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "type": {
                    "description": "Buy / Sell",
                    "type": "integer"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "交易系統",
	Description: "MU-Trading-System",
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
