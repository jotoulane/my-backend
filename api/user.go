package api

import (
	"github.com/gin-gonic/gin"
	"my-backend/global"
)

type UserApi struct {
}

func (pkg *UserApi) UserRegister(c *gin.Context) {
	c.JSON(200, global.ResponseSuccess)
}
