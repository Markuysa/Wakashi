package logger

import (
	"go.uber.org/zap"
)

// InitLogger method creates a logger
func InitLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	config.DisableCaller = true
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
