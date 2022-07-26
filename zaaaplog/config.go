package log

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type Config struct {
	Style            string              // json | console
	WriteSyncer      zapcore.WriteSyncer // *os.File | os.Stderr | os.Stdout
	Level            Level
	TimeFormat       func(t time.Time) string
	ConsoleSeparator string
	Caller           string // fullpath | file | none
	EncodeLevel      string
	LevelBrief       bool
	LevelCapital     bool
	LevelColorful    bool
	ColorMap         ColorMap
}

func (c Config) core() zapcore.Core {
	var enc zapcore.Encoder

	cfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    "function",
		StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     "",
		EncodeLevel:    nil,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration:      nil,
		EncodeCaller:        zapcore.FullCallerEncoder,
		EncodeName:          nil,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    " ",
	}

	if c.TimeFormat != nil {
		cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) { enc.AppendString(c.TimeFormat(t)) }
	}

	if c.ConsoleSeparator != "" {
		cfg.ConsoleSeparator = c.ConsoleSeparator
	}

	switch {
	case c.LevelBrief && c.LevelCapital && c.LevelColorful:
		cfg.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(c.ColorMap._levelToShortCapitalColorString[l])
		}
	case c.LevelBrief && !c.LevelCapital && c.LevelColorful:
		cfg.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(_levelToShortLowercaseColorString[l])
		}
	case c.LevelBrief && c.LevelCapital && !c.LevelColorful:
		cfg.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("[" + l.CapitalString()[0:1] + "]")
		}
	case c.LevelBrief && !c.LevelCapital && !c.LevelColorful:
		cfg.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("[" + l.String()[0:1] + "]")
		}
	case !c.LevelBrief && c.LevelCapital && c.LevelColorful:
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case !c.LevelBrief && !c.LevelCapital && c.LevelColorful:
		cfg.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case !c.LevelBrief && c.LevelCapital && !c.LevelColorful:
		cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	case !c.LevelBrief && !c.LevelCapital && !c.LevelColorful:
		cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	default:
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if c.Caller != "" {
		switch c.Caller {
		case "none":
			cfg.EncodeCaller = nil
		case "file":
			cfg.EncodeCaller = zapcore.ShortCallerEncoder
		case "fullpath":
			cfg.EncodeCaller = zapcore.FullCallerEncoder
		}
		cfg.EncodeCaller = zapcore.FullCallerEncoder
	}

	if c.Style == "json" {
		enc = zapcore.NewJSONEncoder(cfg)
	} else {
		enc = zapcore.NewConsoleEncoder(cfg)
	}
	return zapcore.NewCore(enc, c.WriteSyncer, c.Level)
}
