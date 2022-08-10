package utils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"notification/config"
)

var Logger *zap.Logger

func init() {
	fmt.Println("init logger")

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

	var err error
	Logger, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}
