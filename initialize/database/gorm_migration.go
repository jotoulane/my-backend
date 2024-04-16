package database

import (
	"my-backend/global"
	"my-backend/initialize/logger"
	"my-backend/model/db_model"
)

func migration() {
	err := global.DB.AutoMigrate(
		&db_model.User{},
		//
	)
	if err != nil {
		logger.Logger.Error("model migration failed. %v", err)
	}
}
