package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignUpHandler 注册处理函数
// @Summary 注册功能，用户输入username和password注册
// @Description 注册功能，用户输入username和password注册
// @Tags 注册接口
// @Accept application/json
// @Produce application/json
// @Param LoginParam body models.ParamsSignUp true "注册参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSignUp
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamsSignUp)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("signup err", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 手动校验参数
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 {
	//	zap.L().Error("signup err")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg:": "请求错误",
	//	})
	//}

	fmt.Println(*p)
	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		// 判断错误类型
		if errors.Is(err, mysql.ErrorUserExist) { // 用户名存在
			ResponseError(c, CodeUserExist)
			return
		}
		// 若用户名不存在仍出现错误
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, "注册成功!")
	return
}

// LoginHandler 登录处理
// @Summary 登录功能，用户输入username和password登录，登录成功返回JWT
// @Description 登录功能，用户输入username和password登录，登录成功返回JWT
// @Tags 登录接口
// @Accept application/json
// @Produce application/json
// @Param LoginParam body models.ParamsLogin true "登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseLogin
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	// 1. 获取前端信息(校验密码是否正确)
	user := new(models.ParamsLogin)
	// 获取前端参数错误
	if err := c.ShouldBind(user); err != nil {
		zap.L().Error("login err", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. logic层进行业务处理，获得处理信息
	userWithToken, err := logic.Login(user)
	// 登录失败
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", user.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 3. 登录成功，返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", userWithToken.UserID),
		"user_name": userWithToken.Username,
		"token":     userWithToken.Token,
	})
	return
}
