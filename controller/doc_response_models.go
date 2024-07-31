package controller

import "bluebell/models"

// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}

// _ResponseCreatePost 创建接口响应数据
type _ResponseCreatePost struct {
	Code ResCode `json:"code"` // 业务响应状态码
	Msg  string  `json:"msg"`  // 提示信息
}

type _ResponseLogin struct {
	Code  ResCode `json:"code"`  // 业务响应状态码
	Msg   string  `json:"msg"`   // 提示信息
	Token string  `json:"token"` // 生成的JWT
}

type _ResponseSignUp struct {
	Code ResCode `json:"code"` // 业务响应状态码
	Msg  string  `json:"msg"`  // 提示信息
}
