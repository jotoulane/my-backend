package router

import "github.com/gin-gonic/gin"

type chatRouter struct {
}

func (pkg *chatRouter) InitChatRouter(privateParent *gin.RouterGroup) {
	chatPrivate := privateParent.Group("/chat")
	{
		chatPrivate.POST("send", nil)
	}
}
