package initialize

import (
	"github.com/joho/godotenv"
	"my-backend/initialize/database"
	"my-backend/initialize/logger"
	"os"
)

func Init() {
	// 初始化日志
	logger.InitLogger()

	// 判断是生产环境还是开发环境，并从对应的.env.*文件中读取环境变量。作为判断条件的MODE环境变量在部署环节中由Dockerfile生成镜像时设置
	var err error
	if os.Getenv("MODE") == "PROD" {
		logger.Logger.Info("MODE:", "prod")
		err = godotenv.Load(".env.prod")
	} else {
		logger.Logger.Info("MODE:", "dev")
		err = godotenv.Load(".env.dev")
	}
	if err != nil {
		logger.Logger.Error("Error loading env file", err)
		return
	}

	// 连接数据库
	DATABASE := os.Getenv("DATABASE")
	MYSQL_DSN := os.Getenv("MYSQL_DSN")
	logger.Logger.Info("DATABASE:", DATABASE)
	database.InitGormConnectDB(DATABASE, MYSQL_DSN)

	//初始化数据库内容
	database.InitDBTables()
}
