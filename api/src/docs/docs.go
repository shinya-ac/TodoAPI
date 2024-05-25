// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "localhost:8080",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/task": {
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
                    "Task"
                ],
                "summary": "Task一覧を取得する",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "ページ数",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "ページサイズ",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "Pending",
                        "description": "Todoのステータス",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task.getTaskResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Taskを登録する",
                "parameters": [
                    {
                        "description": "Task登録",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task.CreateTaskParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/task.createTaskResponse"
                        }
                    }
                }
            }
        },
        "/v1/task/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Taskを更新する",
                "parameters": [
                    {
                        "type": "string",
                        "description": "更新するTodoを指定するid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task更新",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task.UpdateTaskParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task.updateTaskResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "task.CreateTaskParams": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "example": "「達人に学ぶクリーンアーキテクチャp100~105」までを読む"
                },
                "title": {
                    "type": "string",
                    "example": "読書"
                }
            }
        },
        "task.Task": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "task.UpdateTaskParams": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 1,
                    "example": "「達人に学ぶクリーンアーキテクチャp200~300」までを読む"
                },
                "status": {
                    "type": "string",
                    "example": "Completed"
                },
                "title": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1,
                    "example": "輪読会"
                }
            }
        },
        "task.createTaskResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "4082ed31-263c-40ec-9d41-e9d274c6bca8"
                }
            }
        },
        "task.getTaskResponse": {
            "type": "object",
            "properties": {
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/task.Task"
                    }
                }
            }
        },
        "task.updateTaskResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "4082ed31-263c-40ec-9d41-e9d274c6bca8"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Todo-API-Key",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Todo API",
	Description:      "RESTful API for TodoApp",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
