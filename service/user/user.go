package user

import (
	"github.com/gin-gonic/gin"
	"my-backend/model/db_model"
)

type UserService struct{}

// Login 登陆
func (pkg *UserService) Login(c *gin.Context, user *db_model.User) error {
	ok, err := user.CheckUser()
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	return err
}
