package logger

import (
	"go.uber.org/zap"
)

func Logger() *zap.Logger {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling:    nil,
	}

	return zap.Must(config.Build())
}
