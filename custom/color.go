package custom

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

// Foreground colors.
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Color represents a text color.
type Color uint8

// Add adds the coloring to the given string.
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}

var (
	_levelToColor = []LevelColor{
		{Level: zapcore.DebugLevel, color: Magenta},
		{Level: zapcore.InfoLevel, color: Blue},
		{Level: zapcore.WarnLevel, color: Yellow},
		{Level: zapcore.ErrorLevel, color: Red},
		{Level: zapcore.DPanicLevel, color: Red},
		{Level: zapcore.PanicLevel, color: Red},
		{Level: zapcore.FatalLevel, color: Red},
	}
	_levelToColorStrings map[string]map[zapcore.Level]string
)

type LevelColor struct {
	zapcore.Level
	color Color
}

func (lc LevelColor) CapitalColorString() string {
	return lc.color.Add(lc.CapitalString())
}

func (lc LevelColor) ColorString() string {
	return lc.color.Add(lc.String())
}

func (lc LevelColor) ShortCapitalColorString() string {
	return lc.color.Add(lc.CapitalString()[0:1])
}

func (lc LevelColor) ShortLowercaseColorString() string {
	return lc.color.Add(lc.String()[0:1])
}

func init() {

}
func FreshLevelToColorStrings(name string) {
	_levelToColorStrings[name] = make(map[zapcore.Level]string, len(_levelToColor))
	for _, levelColor := range _levelToColor {
		_levelToColorStrings[name][levelColor.Level] = levelColor.ColorString()
	}
}

func Set(short, capital, color bool, value string) {

}

func ReplaceLevelToColor(colors map[zapcore.Level]Color) {
	_levelToColor = colors
}
