package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var l = New(os.Stdout, InfoLevel)

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

type Field = zap.Field

type Logger struct {
	l     *zap.Logger
	level Level
}

func (l *Logger) Debug(msg string, fields ...Field)  { l.l.Debug(msg, fields...) }
func (l *Logger) Info(msg string, fields ...Field)   { l.l.Info(msg, fields...) }
func (l *Logger) Warn(msg string, fields ...Field)   { l.l.Warn(msg, fields...) }
func (l *Logger) Error(msg string, fields ...Field)  { l.l.Error(msg, fields...) }
func (l *Logger) DPanic(msg string, fields ...Field) { l.l.DPanic(msg, fields...) }
func (l *Logger) Panic(msg string, fields ...Field)  { l.l.Panic(msg, fields...) }
func (l *Logger) Fatal(msg string, fields ...Field)  { l.l.Fatal(msg, fields...) }

func (l *Logger) Sync() error { return l.l.Sync() }

func Sync() error {
	if l != nil {
		return l.Sync()
	}
	return nil
}

// ResetDefault not safe for concurrent use
func ResetDefault(new *Logger) {
	l = new
}

func New(writer io.Writer, level Level, opts ...zap.Option) *Logger {
	if writer == nil {
		panic("log writer is nil")
	}
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		level,
	)
	return &Logger{
		l:     zap.New(core, opts...),
		level: level,
	}
}

func Debug(msg string, fields ...Field)  { l.l.Debug(msg, fields...) }
func Info(msg string, fields ...Field)   { l.l.Info(msg, fields...) }
func Warn(msg string, fields ...Field)   { l.l.Warn(msg, fields...) }
func Error(msg string, fields ...Field)  { l.l.Error(msg, fields...) }
func DPanic(msg string, fields ...Field) { l.l.DPanic(msg, fields...) }
func Panic(msg string, fields ...Field)  { l.l.Panic(msg, fields...) }
func Fatal(msg string, fields ...Field)  { l.l.Fatal(msg, fields...) }
