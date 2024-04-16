package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"my-backend/global"
	log2 "my-backend/initialize/logger"
	"os"
	"time"
)

func InitGormConnectDB(DATABASE, DSN string) {
	// 初始化GORM日志配置
	var logWriter io.Writer = *log2.LogFileWriter
	newLogger := logger.New(
		log.New(io.MultiWriter(logWriter, os.Stdout), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	var dialector gorm.Dialector
	dialector = mysql.Open(DSN)

	//连接数据库
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true, // 迁移时不要创建外键约束
	})

	if err != nil {
		log2.Logger.Error("database lost: %v", err)
		panic(interface{}(err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		log2.Logger.Error("database lost: %v", err)
		panic(interface{}(err))
	}

	//设置连接池
	//设置数据库连接池中的最大空闲连接数为10
	sqlDB.SetMaxIdleConns(10)
	//设置数据库连接池中的最大打开连接数为20
	sqlDB.SetMaxOpenConns(20)
	global.DB = db
	global.DBWithoutLog = db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})

	migration()
}
