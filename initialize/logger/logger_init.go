package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var (
	Logger        *zap.SugaredLogger
	LogFileWriter *zapcore.WriteSyncer
)

// getLogWriter 保存文件日志切割
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/ia.log", // 指定日志文件目录
		MaxSize:    20,              // 文件内容大小, MB
		MaxBackups: 5,               // 保留旧文件最大个数
		MaxAge:     30,              // 保留旧文件最大天数
		Compress:   false,           // 文件是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getEncoder 获取日志编码格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 格式化时间 可自定义
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zap.CombineWriteSyncers(writeSyncer), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Logger = logger.Sugar()

	// 设置gin内置log
	gin.DefaultWriter = io.MultiWriter(writeSyncer, os.Stdout)

	// 设置gorm内置log
	LogFileWriter = &writeSyncer

	Logger.Info("Logger init ok!")
}
