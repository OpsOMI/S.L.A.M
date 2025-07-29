package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLogger returns a zap logger that always logs to console.
func NewZapLogger() (ILogger, error) {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.InfoLevel)

	logger := zap.New(consoleCore, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.FatalLevel))

	return &zapLogger{
		logger:        logger,
		sugaredLogger: logger.Sugar(),
	}, nil
}

func (l *zapLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *zapLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *zapLogger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *zapLogger) Infof(template string, args ...any) {
	l.sugaredLogger.Infof(template, args...)
}

func (l *zapLogger) Errorf(template string, args ...any) {
	l.sugaredLogger.Errorf(template, args...)
}

func (l *zapLogger) Debugf(template string, args ...any) {
	l.sugaredLogger.Debugf(template, args...)
}

func (l *zapLogger) Warnf(template string, args ...any) {
	l.sugaredLogger.Warnf(template, args...)
}

func (l *zapLogger) Sync() error {
	return l.logger.Sync()
}
