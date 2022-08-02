package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"notification/config"
)

var Logger *zap.Logger

func InitLog() (*zap.Logger, error) {
	var atomicLevel zapcore.Level
	if config.Config.Debug {
		atomicLevel = zapcore.DebugLevel
	} else {
		atomicLevel = zapcore.InfoLevel
	}
	logConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(atomicLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return logConfig.Build()
}

func MyLog(model string, action string, objectID, userID int, msg ...string) {
	message := ""
	for _, v := range msg {
		message += v
	}
	Logger.Info(
		model+""+action,
		zap.Int("UserID", userID),
		zap.Int("ID", objectID),
		zap.String("Additional", message),
	)
}
