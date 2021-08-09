/*
 * @Author: Ali2vu <751815097@qq.com>
 * @Date: 2021-08-09 23:54:32
 * @LastEditors: Ali2vu
 * @LastEditTime: 2021-08-10 01:25:06
 */
package logger

import (
	"fastmini/config"
	"fastmini/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	if config.Config("APP_ENV", "dev") == "dev" {
		logger, _ := zap.NewDevelopment()
		global.FM_LOG = logger.Sugar()

	} else {
		writeSyncer := getLogWriter()
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
		logger := zap.New(core, zap.AddCaller())
		global.FM_LOG = logger.Sugar()
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./app.log",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
