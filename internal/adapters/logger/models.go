package logger

import "go.uber.org/zap"

// ILogger interface defines logging methods.
type ILogger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Sync() error

	// Sugared methods for easier logging without fields
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
}

type zapLogger struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}
