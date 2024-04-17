package api

import (
	"github.com/gin-gonic/gin"
	"my-backend/global/code"
	"my-backend/global/response"
	"my-backend/initialize/logger"
	"my-backend/middleware"
	"my-backend/model/db_model"
	"os"
)

type UserApi struct {
}

func (pkg *UserApi) Ping(c *gin.Context) {
	logger.Logger.Info("ping")
	c.JSON(code.StatusOK, response.ResponseSuccess("pong"))
}

func (pkg *UserApi) Login(c *gin.Context) {
	//参数验证
	var form struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required,min=5,max=40"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(code.StatusOK, response.ResponseError(code.StatusBadRequest))
		return
	}
	//业务逻辑
	user := db_model.User{
		Phone:    form.Phone,
		Password: form.Password,
	}
	err := userService.Login(c, &user)
	if err != nil {
		c.JSON(200, response.ResponseErrorWithMsg(code.StatusBadRequest, err.Error()))
		return
	}

	// 登录成功，生成JWT
	token, err := middleware.GenToken(user.ID, user.UserName, user.Phone)
	if err != nil {
		c.JSON(200, response.ResponseErrorWithMsg(code.StatusInternalServerError, err.Error()))
		return
	}

	// 回传token和用户基本信息
	c.JSON(200, response.ResponseSuccess(gin.H{
		"token":    token,
		"userInfo": user,
	}))
}

func (pkg *UserApi) UserRegister(c *gin.Context) {
	c.JSON(200, response.ResponseSuccess(os.Getenv("DATABASE")))
}
