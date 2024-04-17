package router

import (
	"github.com/gin-gonic/gin"
	"my-backend/api"
)

type userRouter struct {
}

func (pkg *userRouter) InitUserRouter(privateParent *gin.RouterGroup, publicParent *gin.RouterGroup) {
	userPublic := publicParent.Group("/user")
	//userPrivate := privateParent.Group("/user")

	userApi := api.ApiPackageApp.UserApi

	// 公共接口注册
	{
		userPublic.GET("/ping", userApi.Ping)
		userPublic.POST("/login", userApi.Login)
		userPublic.GET("/register", userApi.UserRegister)
	}

}
