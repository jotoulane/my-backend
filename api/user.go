package api

import (
	"github.com/gin-gonic/gin"
	"my-backend/model/code"
	"my-backend/model/response"
)

type UserApi struct {
}

func (pkg *UserApi) Ping(c *gin.Context) {
	c.JSON(code.StatusOK, response.ResponseSuccess("pong"))
}

func (pkg *UserApi) UserRegister(c *gin.Context) {
	c.JSON(200, response.ResponseSuccess("register success"))
}
