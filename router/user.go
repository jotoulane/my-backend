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

	{
		userPublic.GET("/ping", userApi.Ping)
		userPublic.GET("/register", userApi.UserRegister)
	}

}
