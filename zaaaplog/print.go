package log

func Debug(msg string, fields ...Field)  { l.Debug(msg, fields...) }
func Info(msg string, fields ...Field)   { l.Info(msg, fields...) }
func Warn(msg string, fields ...Field)   { l.Warn(msg, fields...) }
func Error(msg string, fields ...Field)  { l.Error(msg, fields...) }
func DPanic(msg string, fields ...Field) { l.DPanic(msg, fields...) }
func Panic(msg string, fields ...Field)  { l.Panic(msg, fields...) }
func Fatal(msg string, fields ...Field)  { l.Fatal(msg, fields...) }

func Debugf(msg string, args ...interface{})  { l.Sugar().Debugf(msg, args...) }
func Infof(msg string, args ...interface{})   { l.Sugar().Infof(msg, args...) }
func Warnf(msg string, args ...interface{})   { l.Sugar().Warnf(msg, args...) }
func Errorf(msg string, args ...interface{})  { l.Sugar().Errorf(msg, args...) }
func DPanicf(msg string, args ...interface{}) { l.Sugar().DPanicf(msg, args...) }
func Panicf(msg string, args ...interface{})  { l.Sugar().Panicf(msg, args...) }
func Fatalf(msg string, args ...interface{})  { l.Sugar().Fatalf(msg, args...) }
