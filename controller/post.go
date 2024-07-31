package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子接口
// @Summary 创建帖子接口
// @Description 根据前端传递的参数创建帖子，存到mysql数据库中
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT的格式"
// @Param object body models.ParamNewPost true "帖子信息"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCreatePost
// @Router /createPost [get]
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数的检验 c.ShouldBindJSON(p)
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) error: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取当前用户id
	authorID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
	}
	p.AuthorID = authorID
	// 2. 创建贴子 logic.CreatePost()
	if err := logic.CreatePost(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, "创建Post成功")
}

// GetPostDetailHandler
// @Summary 根据某个帖子id获取单个帖子详细信息
// @Description 根据某个帖子id获取单个帖子详细信息
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param postID query int64 false "帖子ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /getPost/:id [get]
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取url 中的参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 根据id从数据库中查询数据
	postData, err := logic.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID error ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 3. 返回响应，通过json格式将数据返回
	ResponseSuccess(c, *postData)
}

// GetPostListHandler 获取帖子列表接口
// @Summary 获取帖子列表接口
// @Description 根据前端传递的(分页)参数按时间或分数排序查询帖子列表接口，mysql
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param Page query int64 false "页面数量"
// @Param Size query int64 false "页面大小"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /getPosts [get]
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size, err := getPageInfo(c)
	// 1. 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() error ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 2. 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 根据前端传递的(分页)参数按时间或分数排序查询帖子列表接口，redis
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /getPosts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// 1. 获取分页参数 ShouldBind
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
		//CommunityID: 1, // 可以为空
	}
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("GetPostListHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if p.Order != models.OrderTime && p.Order != models.OrderScore {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 获取数据
	data, err := logic.GetPostListNew(p) // 合二为一(community_id是否为空)
	if err != nil {
		zap.L().Error("logic.GetPostListNew() error ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, data)
	return
}

// GetPostListByCommunityHandler
// 根据社区id查询帖子列表
//func GetPostListByCommunityHandler(c *gin.Context) {
//	// 1. 获取分页参数 ShouldBind
//	post := &models.ParamCommunityPostList{
//		ParamPostList: &models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//		CommunityID: 1,
//	}
//	if err := c.ShouldBind(&post); err != nil {
//		zap.L().Error("GetPostListHandler with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//	zap.L().Debug("p.community_id = ", zap.String("community_id", strconv.FormatInt(post.CommunityID, 10)))
//	if post.Order != models.OrderTime && post.Order != models.OrderScore {
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//	// 2. 获取数据
//	data, err := logic.GetCommunityPostList(post)
//	if err != nil {
//		zap.L().Error("logic.GetPostList2() error ", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//
//	// 3. 返回响应
//	ResponseSuccess(c, data)
//	return
//}
