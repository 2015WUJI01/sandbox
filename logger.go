package main

import (
	"fmt"
	gologger "log"
)

/*

颜色	前景色	背景色
黑色	\033[30m	\033[40m
红色	\033[31m	\033[41m
绿色	\033[32m	\033[42m
橙色	\033[33m	\033[43m
蓝色	\033[34m	\033[44m
品红	\033[35m	\033[45m
青色	\033[36m	\033[46m
浅灰	\033[37m	\033[47m
回退到发行版默认值	\033[39m	\033[49m

颜色	背景色
深灰	\033[100m
浅红	\033[101m
浅绿	\033[102m
黄色	\033[103m
浅蓝	\033[104m
浅紫	\033[105m
蓝绿	\033[106m
白色	\033[107m
*/
const (
	origin = "\033[39m"
	blue   = "\033[34m"
	warn   = "\033[33m"
	red    = "\033[35m"
)

var log Logger

type Logger struct{}

func init() {
	log = Logger{}
}

func (l Logger) Infof(msg string, args ...interface{}) {
	logtype := fmt.Sprintf("%s%-8s%s", blue, "[INFO]", origin)
	gologger.Printf(logtype+msg, args...)
}

func (l Logger) Warnf(msg string, args ...interface{}) {
	logtype := fmt.Sprintf("%s%-8s%s", warn, "[Warn]", origin)
	gologger.Printf(logtype+msg, args...)
}

func (l Logger) Errorf(msg string, args ...interface{}) {
	logtype := fmt.Sprintf("%s%-8s%s", red, "[ERROR]", origin)
	gologger.Printf(logtype+msg, args...)
}
