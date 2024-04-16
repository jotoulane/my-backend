package database

import (
	"my-backend/global"
	"my-backend/initialize/logger"
)

func migration() {
	err := global.DB.AutoMigrate(
	//
	)
	if err != nil {
		logger.Logger.Error("model migration failed. %v", err)
	}
}
