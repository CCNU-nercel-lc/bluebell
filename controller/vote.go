package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PostVoteHandler 投票处理函数
func PostVoteHandler(c *gin.Context) {
	// 参数校验 用ShouldBindJSON
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		//_, ok := err.(validator.ValidationErrors) // 类型断言，看是否触发ShouldBindJSON错误
		//if !ok {
		//	ResponseError(c, CodeInvalidParam)
		//	fmt.Printf("ResponseError\n")
		//	return
		//}
		zap.L().Error("c.ShouldBindJSON(p) failed ", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "参数错误")
		return
	}

	// 获取当前请求的用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	// 调用logic投票
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost(userID, p) failed", zap.Error(err))
		ResponseError(c, CodeRepeatVote)
		return
	}
	ResponseSuccess(c, nil)

}
