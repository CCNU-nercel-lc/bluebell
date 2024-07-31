package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(mode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 分组v1
	v1 := r.Group("/api/v1")

	// test 路由
	//r.POST("/test", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	// 从请求头中的Token判断是否登录
	//	controller.ResponseSuccess(c, "test success")
	//
	//})

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	// 登录业务路由
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.GET("/getPost/:id", controller.GetPostDetailHandler)
	v1.GET("/getPosts", controller.GetPostListHandler)
	v1.GET("/getPosts2", controller.GetPostListHandler2)

	// JWT验证中间件
	//v1.Use(middlewares.RateLimitMiddleware(2*time.Second, 1), middlewares.JWTAuthMiddleware())
	v1.Use(middlewares.JWTAuthMiddleware())

	{

		v1.POST("/createPost", controller.CreatePostHandler)

		//v1.GET("/getlistbycommunity", controller.GetPostListByCommunityHandler)
		v1.POST("/vote", controller.PostVoteHandler)

	}

	pprof.Register(r) // pprof

	return r
}
