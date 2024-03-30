package initialize

import (
	"github.com/gin-gonic/gin"
	"my-backend/middleware"
	"my-backend/router"
)

func InitRoutes() *gin.Engine {
	r := gin.New()

	// 生成路由引擎
	v1Publish := r.Group("/api/v1")
	v1Private := r.Group("/api/v1")
	v1Private.Use(middleware.JWTAuthMiddleware())

	// 初始化路由
	routerApp := router.RouterPackageApp

	routerApp.UserRouter.InitUserRouter(v1Private, v1Publish)
	routerApp.ChatRouter.InitChatRouter(v1Private)

	return r
}
