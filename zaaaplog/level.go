package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	DebugLevel  = zap.DebugLevel  // -1
	InfoLevel   = zap.InfoLevel   // 0, default
	WarnLevel   = zap.WarnLevel   // 1
	ErrorLevel  = zap.ErrorLevel  // 2
	DPanicLevel = zap.DPanicLevel // 3, used in development log
	PanicLevel  = zap.PanicLevel  // 4, logs a message, then panics
	FatalLevel  = zap.FatalLevel  // 5, logs a message, then calls os.Exit(1)
)
