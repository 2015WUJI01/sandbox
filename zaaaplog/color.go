package log

import (
	"fmt"
	"strconv"
	"strings"
)

type Color interface {
	SetForeground(s string) string
	SetBackground(s string) string
}

type ANSI uint8

func (ansi ANSI) SetForeground(s string) string {
	return fmt.Sprintf("\x1b[1;%d;%dm%s\x1b[0m", _fgc, ansi, s)
}
func (ansi ANSI) SetBackground(s string) string {
	return fmt.Sprintf("\x1b[1;%d;%dm%s\x1b[0m", _bgc, ansi, s)
}

type RGB struct {
	r, g, b uint32
	hex     string
}

func (c RGB) SetForeground(s string) string {
	return fmt.Sprintf("\x1b[1;%d;%d;%d;%d;%dm%s\x1b[0m", _fgc, _rgb, c.r, c.g, c.b, s)
}

func (c RGB) SetBackground(s string) string {
	return fmt.Sprintf("\x1b[1;%d;%d;%d;%d;%dm%s\x1b[0m", _bgc, _rgb, c.r, c.g, c.b, s)
}

func NewRGBwithHex(hex string) RGB {
	s := strings.TrimLeft(strings.TrimSpace(hex), "#")
	if len(s) == 3 || len(s) == 6 {
		for _, char := range s {
			switch {
			case char >= '0' && char <= '9':
			case char >= 'A' && char <= 'F':
			case char >= 'a' && char <= 'f':
			default:
				Error("invalid hex value")
				return RGB{}
			}
		}
		if len(s) == 3 {
			s = s[0:1] + s[0:1] + s[1:2] + s[1:2] + s[2:3] + s[2:3]
		}
		r, _ := hex2Uint32(s[0:2])
		g, _ := hex2Uint32(s[2:4])
		b, _ := hex2Uint32(s[4:6])
		return RGB{r: r, g: g, b: b, hex: "#" + s}
	}
	return RGB{}
}
func NewRGBwithUint32(r, g, b uint32) RGB {
	return RGB{r: r, g: g, b: b, hex: fmt.Sprintf("#%02x%02x%02x", r, g, b)}
}

const (
	BlackANSI ANSI = iota + 30
	RedANSI
	GreenANSI
	YellowANSI
	BlueANSI
	MagentaANSI
	CyanANSI
	WhiteANSI
)

var (
	Black   = NewRGBwithHex("#000")
	Red     = NewRGBwithUint32(205, 49, 49)
	Green   = NewRGBwithUint32(13, 188, 121)
	Yellow  = NewRGBwithUint32(229, 229, 16)
	Blue    = NewRGBwithUint32(36, 114, 200)
	Magenta = NewRGBwithUint32(188, 63, 188)
	Cyan    = NewRGBwithUint32(17, 168, 205)
	White   = NewRGBwithUint32(229, 229, 229)

	BrightBlack   = NewRGBwithUint32(85, 85, 85)
	BrightRed     = NewRGBwithUint32(255, 85, 85)
	BrightGreen   = NewRGBwithUint32(35, 209, 139)
	BrightYellow  = NewRGBwithUint32(245, 245, 67)
	BrightBlue    = NewRGBwithUint32(59, 142, 234)
	BrightMagenta = NewRGBwithUint32(188, 63, 188)
	BrightCyan    = NewRGBwithUint32(41, 184, 219)
	BrightWhite   = NewRGBwithHex("#fff")
)

func hex2Uint32(s string) (uint32, error) {
	i, err := strconv.ParseUint(s, 16, 32)
	return uint32(i), err
}

const (
	_fgc = 38
	_bgc = 48
	_rgb = 2
)

var (
	// 定义一组默认的颜色
	_levelToColor = DefaultColorMap()

	_unknownLevelColor = Red

	_levelToShortLowercaseColorString = make(map[Level]string, len(_levelToColor))
	_levelToShortCapitalColorString   = make(map[Level]string, len(_levelToColor))
)

func init() {
	for level, color := range _levelToColor {
		_levelToShortLowercaseColorString[level] = color.SetForeground("[" + level.CapitalString()[0:1] + "]")
		_levelToShortCapitalColorString[level] = color.SetForeground("[" + level.CapitalString()[0:1] + "]")
	}
}

func DefaultColorMap() map[Level]Color {
	return map[Level]Color{
		DebugLevel:  GreenANSI,
		InfoLevel:   BlueANSI,
		WarnLevel:   YellowANSI,
		ErrorLevel:  RedANSI,
		DPanicLevel: RedANSI,
		PanicLevel:  RedANSI,
		FatalLevel:  RedANSI,
	}
}

type LevelColor struct {
	Debug  Color
	Info   Color
	Warn   Color
	Error  Color
	DPanic Color
	Panic  Color
	Fatal  Color
}
