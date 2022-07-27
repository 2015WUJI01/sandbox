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

	LogLevel struct {
		Brief    bool
		Capital  bool
		Colorful bool
	}
	LevelBrief    bool
	LevelCapital  bool
	LevelColorful bool

	LevelColor                       LevelColor
	levelToBriefLowercaseColorString map[Level]string
	levelToBriefCapitalColorString   map[Level]string
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

	levelColorMap := DefaultColorMap()
	if c.LevelColor.Debug == nil {
		c.LevelColor.Debug = levelColorMap[DebugLevel]
	}
	if c.LevelColor.Info == nil {
		c.LevelColor.Info = levelColorMap[InfoLevel]
	}
	if c.LevelColor.Warn == nil {
		c.LevelColor.Warn = levelColorMap[WarnLevel]
	}
	if c.LevelColor.Error == nil {
		c.LevelColor.Error = levelColorMap[ErrorLevel]
	}
	if c.LevelColor.Panic == nil {
		c.LevelColor.Panic = levelColorMap[PanicLevel]
	}
	if c.LevelColor.DPanic == nil {
		c.LevelColor.DPanic = levelColorMap[DPanicLevel]
	}
	if c.LevelColor.Fatal == nil {
		c.LevelColor.Fatal = levelColorMap[FatalLevel]
	}

	c.levelToBriefLowercaseColorString = func() map[Level]string {
		levelstring := make(map[Level]string, 7)
		levelstring[DebugLevel] = c.LevelColor.Debug.SetForeground("[" + DebugLevel.String()[0:1] + "]")
		levelstring[InfoLevel] = c.LevelColor.Info.SetForeground("[" + InfoLevel.String()[0:1] + "]")
		levelstring[WarnLevel] = c.LevelColor.Warn.SetForeground("[" + WarnLevel.String()[0:1] + "]")
		levelstring[ErrorLevel] = c.LevelColor.Error.SetForeground("[" + ErrorLevel.String()[0:1] + "]")
		levelstring[PanicLevel] = c.LevelColor.Panic.SetForeground("[" + PanicLevel.String()[0:1] + "]")
		levelstring[DPanicLevel] = c.LevelColor.DPanic.SetForeground("[" + DPanicLevel.String()[0:1] + "]")
		levelstring[FatalLevel] = c.LevelColor.Fatal.SetForeground("[" + FatalLevel.String()[0:1] + "]")
		return levelstring
	}()

	c.levelToBriefCapitalColorString = func() map[Level]string {
		levelstring := make(map[Level]string, 7)
		levelstring[DebugLevel] = c.LevelColor.Debug.SetForeground("[" + DebugLevel.CapitalString()[0:1] + "]")
		levelstring[InfoLevel] = c.LevelColor.Info.SetForeground("[" + InfoLevel.CapitalString()[0:1] + "]")
		levelstring[WarnLevel] = c.LevelColor.Warn.SetForeground("[" + WarnLevel.CapitalString()[0:1] + "]")
		levelstring[ErrorLevel] = c.LevelColor.Error.SetForeground("[" + ErrorLevel.CapitalString()[0:1] + "]")
		levelstring[PanicLevel] = c.LevelColor.Panic.SetForeground("[" + PanicLevel.CapitalString()[0:1] + "]")
		levelstring[DPanicLevel] = c.LevelColor.DPanic.SetForeground("[" + DPanicLevel.CapitalString()[0:1] + "]")
		levelstring[FatalLevel] = c.LevelColor.Fatal.SetForeground("[" + FatalLevel.CapitalString()[0:1] + "]")
		return levelstring
	}()

	switch {
	case c.LevelBrief && c.LevelCapital && c.LevelColorful:
		cfg.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(c.levelToBriefCapitalColorString[l])
		}
	case c.LevelBrief && !c.LevelCapital && c.LevelColorful:
		cfg.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(c.levelToBriefLowercaseColorString[l])
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
