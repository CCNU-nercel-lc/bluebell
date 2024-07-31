// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "这里写联系人信息",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/createPost": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据前端传递的参数创建帖子，存到mysql数据库中",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "创建帖子接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT的格式",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "帖子信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamNewPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseCreatePost"
                        }
                    }
                }
            }
        },
        "/getPost/:id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据某个帖子id获取单个帖子详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "根据某个帖子id获取单个帖子详细信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "帖子ID",
                        "name": "postID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        },
        "/getPosts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据前端传递的(分页)参数按时间或分数排序查询帖子列表接口，mysql",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "获取帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页面数量",
                        "name": "Page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页面大小",
                        "name": "Size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        },
        "/getPosts2": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据前端传递的(分页)参数按时间或分数排序查询帖子列表接口，redis",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "升级版帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "可以为空",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序依据(time or score)",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页面数量",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页面大小",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "登录功能，用户输入username和password登录，登录成功返回JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录接口"
                ],
                "summary": "登录功能，用户输入username和password登录，登录成功返回JWT",
                "parameters": [
                    {
                        "description": "登录参数",
                        "name": "LoginParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamsLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseLogin"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "注册功能，用户输入username和password注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "注册接口"
                ],
                "summary": "注册功能，用户输入username和password注册",
                "parameters": [
                    {
                        "description": "注册参数",
                        "name": "LoginParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamsSignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseSignUp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ResCode": {
            "type": "integer",
            "enum": [
                1000,
                1001,
                1002,
                1003,
                1004,
                1005,
                1006,
                1007
            ],
            "x-enum-varnames": [
                "CodeSuccess",
                "CodeInvalidParam",
                "CodeUserExist",
                "CodeUserNotExist",
                "CodeInvalidPassword",
                "CodeServerBusy",
                "CodeNotLogin",
                "CodeInvalidToken"
            ]
        },
        "controller._ResponseCreatePost": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "controller._ResponseLogin": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                },
                "token": {
                    "description": "生成的JWT",
                    "type": "string"
                }
            }
        },
        "controller._ResponsePostList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ApiPostDetail"
                    }
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "controller._ResponseSignUp": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models.ApiPostDetail": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "CommunityDetail": {
                    "description": "嵌入社区信息结构体",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.CommunityDetail"
                        }
                    ]
                },
                "author_id": {
                    "description": "用户ID，登录时保存到JWT中，通过解析JWT将ID放入ctx中",
                    "type": "string",
                    "example": "0"
                },
                "author_name": {
                    "type": "string"
                },
                "community_id": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "description": "帖子id，通过雪花算法生成",
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "vote_num": {
                    "type": "integer"
                }
            }
        },
        "models.CommunityDetail": {
            "type": "object",
            "properties": {
                "create_Time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.ParamNewPost": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "community_id": {
                    "description": "所属社区id",
                    "type": "integer"
                },
                "content": {
                    "description": "内容",
                    "type": "string"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                }
            }
        },
        "models.ParamsLogin": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "models.ParamsSignUp": {
            "type": "object",
            "required": [
                "password",
                "rePassword",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "rePassword": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "天涯论坛2024版",
	Description:      "这里写描述信息",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}