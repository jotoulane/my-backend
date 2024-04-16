package database

import (
	"golang.org/x/crypto/bcrypt"
	"my-backend/global"
	"my-backend/initialize/logger"
	"my-backend/model/db_model"
)

func InitDBTables() {
	var db = global.DB
	logger.Logger.Info("初始化数据库表")

	// 检查数据库中的用户表是否为空
	var userCount int64
	db.Table("users").Count(&userCount)
	if userCount == 0 {
		logger.Logger.Info("用户表为空，开始初始化")
		password, err2 := bcrypt.GenerateFromPassword([]byte(db_model.DefaultPassword), db_model.PassWordCost)
		if err2 != nil {
			logger.Logger.Error("初始化用户表失败", err2)
		}
		// 初始化用户表
		superUser := db_model.User{
			UserName: "admin",
			Phone:    "12345678910",
			Password: string(password),
		}
		err := superUser.Insert()
		if err != nil {
			logger.Logger.Error("初始化用户表失败", err)
		} else {
			logger.Logger.Info("初始化用户表成功")
		}
	}
}
