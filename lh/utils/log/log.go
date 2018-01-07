package log

import (
	"os"
)

var defaultLog Log = nil

func init() {
	SetDriver(CliLog{})
}

func SetDriver(destLog Log) {
	defaultLog = destLog
}

func DisableLog() {
	defaultLog = DisabledLog{}
}

func Debug(v ...interface{}) {
	defaultLog.Output(LL_DEBUG, v...)
}
func Debugf(format string, v ...interface{}) {
	defaultLog.Outputf(LL_DEBUG, format, v...)
}
func Info(v ...interface{}) {
	defaultLog.Output(LL_INFO, v...)
}
func Infof(format string, v ...interface{}) {
	defaultLog.Outputf(LL_INFO, format, v...)
}
func Warn(v ...interface{}) {
	defaultLog.Output(LL_WARNING, v...)
}
func Warnf(format string, v ...interface{}) {
	defaultLog.Outputf(LL_WARNING, format, v...)
}
func Error(v ...interface{}) {
	defaultLog.Output(LL_ERROR, v...)
}
func Errorf(format string, v ...interface{}) {
	defaultLog.Outputf(LL_ERROR, format, v...)
}
func Fatal(code int, v ...interface{}) {
	defaultLog.Output(LL_CRITICAL, v...)
	os.Exit(code)
}
func Fatalf(code int, format string, v ...interface{}) {
	defaultLog.Outputf(LL_CRITICAL, format, v...)
	os.Exit(code)
}
