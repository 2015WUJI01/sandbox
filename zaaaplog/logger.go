package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var l = DefaultLogger()

type Logger struct {
	*zap.Logger
}

func DefaultLogger() *Logger {
	return NewLogger(Config{
		Level: DebugLevel,
	})
}

func NewLogger(cfg Config) *Logger {
	return &Logger{
		Logger: zap.New(cfg.core()),
	}
}

func ResetLogger(new *Logger) {
	l = new
}

func newcore() zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		DebugLevel,
	)
}
