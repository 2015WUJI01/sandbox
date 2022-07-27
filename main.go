package main

import (
	log "main/zaaaplog"
	"os"
	"time"
)

func main() {

	log.ResetLogger(log.NewLogger(log.Config{

		// print style and output path
		Style:       "console",
		WriteSyncer: os.Stdout,

		// min logging level
		Level: log.DebugLevel,

		// logging formatter

		// Separator format
		Separator: " ", // only useful in console style

		// time format
		TimeFormat: func(t time.Time) string {
			return ""
		},

		// level format
		LevelFormat: func(l log.Level, c log.Color) string {
			return ""
		},

		// colorful
		LevelColor: log.LevelColor{
			Info:   log.BrightBlue,
			Warn:   log.YellowANSI,
			Error:  log.RedANSI,
			DPanic: log.RedANSI,
			Panic:  log.RedANSI,
			Fatal:  log.RedANSI,
		},
	}))

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.DPanic("dpanic")
	func() {
		defer func() {
			recover()
		}()
		log.Panic("panic")
	}()
	log.Fatal("fatal")
}

//
// import (
// 	"fmt"
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// 	"io"
// 	"os"
// 	"time"
// )
//
// // RGBColor represents a text color.
// type RGBColor uint8
//
// // Add adds the coloring to the given string.
// func (c RGBColor) Add(s string) string {
// 	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
// }
//
// type style int8
//
// const (
// 	JsonStyle style = iota
// 	ConsoleStyle
// )
//
// var (
// 	_levelToColor = map[zapcore.Level]RGBColor{
// 		zapcore.DebugLevel:  Magenta,
// 		zapcore.InfoLevel:   Blue,
// 		zapcore.WarnLevel:   Yellow,
// 		zapcore.ErrorLevel:  Red,
// 		zapcore.DPanicLevel: Red,
// 		zapcore.PanicLevel:  Red,
// 		zapcore.FatalLevel:  Red,
// 	}
// 	_unknownLevelColor = Red
//
// 	_levelToLowercaseColorString      = make(map[zapcore.Level]string, len(_levelToColor))
// 	_levelToCapitalColorString        = make(map[zapcore.Level]string, len(_levelToColor))
// 	_levelToShortLowercaseColorString = make(map[zapcore.Level]string, len(_levelToColor))
// 	_levelToShortCapitalColorString   = make(map[zapcore.Level]string, len(_levelToColor))
// )
//
// func init() {
// 	for level, color := range _levelToColor {
// 		_levelToLowercaseColorString[level] = color.Add(level.String())
// 		_levelToCapitalColorString[level] = color.Add(level.CapitalString())
// 		_levelToShortLowercaseColorString[level] = color.Add(level.CapitalString()[0:1])
// 		_levelToShortCapitalColorString[level] = color.Add(level.CapitalString()[0:1])
// 	}
// }
// func main() {
// 	log := zap.New(zapcore.NewTee(
// 		core(os.Stdout, ConsoleStyle),
// 		core(getWriter("app.log"), ConsoleStyle),
// 		core(getWriter("log.json"), JsonStyle),
// 	), zap.AddCaller())
// 	for range time.Tick(500 * time.Millisecond) {
// 		log.Info("啦啦啦")
// 	}
// }
//
// func core(w io.Writer, s style) zapcore.Core {
// 	cfg := zap.NewProductionEncoderConfig()
// 	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 		enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
// 	}
//
// 	var colors = map[zapcore.Level]int8{
// 		zapcore.DebugLevel: 35, zapcore.InfoLevel: 34, zapcore.WarnLevel: 33, zapcore.ErrorLevel: 31,
// 		zapcore.DPanicLevel: 31, zapcore.PanicLevel: 31, zapcore.FatalLevel: 31,
// 	}
// 	cfg.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
// 		enc.AppendString(fmt.Sprintf("\x1b[%dm[%s]\x1b[0m", colors[l], l.CapitalString()[0:1]))
// 	}
// 	if s == JsonStyle {
// 		return zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(w), zapcore.DebugLevel)
// 	}
// 	cfg.ConsoleSeparator = " "
// 	cfg.EncodeCaller = FullPathCallerEncoder
// 	cfg.EncodeName = func(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
// 		enc.AppendString(loggerName)
// 	}
// 	return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.AddSync(w), zapcore.DebugLevel)
// }
//
// func getWriter(path string) io.Writer {
// 	file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
// 	return file
// }
//
// func NewEncoderConfig() zapcore.EncoderConfig {
// 	cfg := zap.NewProductionEncoderConfig()
// 	cfg.EncodeCaller = FullPathCallerEncoder
// 	cfg.EncodeTime = LocalTimeEncoder
// 	return cfg
// }
//
// func NewJsonEncoderConfig() zapcore.EncoderConfig {
// 	cfg := NewEncoderConfig()
// 	cfg.EncodeCaller = FullPathCallerEncoder
// 	cfg.EncodeTime = LocalTimeEncoder
// 	return cfg
// }
// func NewConsoleEncoderConfig() zapcore.EncoderConfig {
// 	cfg := NewEncoderConfig()
// 	return cfg
// }
//
// func ShortCapitalColorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
// 	var colors = map[zapcore.Level]int8{
// 		zapcore.DebugLevel: 35, zapcore.InfoLevel: 34, zapcore.WarnLevel: 33, zapcore.ErrorLevel: 31,
// 		zapcore.DPanicLevel: 31, zapcore.PanicLevel: 31, zapcore.FatalLevel: 31,
// 	}
// 	// 去 map 中读取
// 	enc.AppendString(fmt.Sprintf("\x1b[%dm[%s]\x1b[0m", colors[l], l.CapitalString()[0:1]))
// }
//
// func LocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
// }
//
// func FullPathCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString(caller.FullPath())
// }
//
// // Foreground colors.
// const (
// 	Black RGBColor = iota + 30
// 	Red
// 	Green
// 	Yellow
// 	Blue
// 	Magenta
// 	Cyan
// 	White
// )
