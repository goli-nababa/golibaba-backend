package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  *zap.Logger
	once sync.Once
)

func NewLogger(level zapcore.Level, isProduction bool) {
	once.Do(func() {
		var config zap.Config
		if isProduction {
			config = zap.NewProductionConfig()
		} else {
			config = zap.NewDevelopmentConfig()
		}

		config.Level.SetLevel(level)

		var err error

		log, err = config.Build()
		if err != nil {
			panic(err)
		}

	})
}

func GetLogger() *zap.Logger {
	if log == nil {
		NewLogger(zap.InfoLevel, true)
	}
	return log
}

func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}
