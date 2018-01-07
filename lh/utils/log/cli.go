package log

import (
	"fmt"
	"time"

	"github.com/467754239/GolangNote/lh/config"
)

type CliPerttyLog struct {
}
type CliLog struct {
}

func (this CliLog) Output(level LogLevel, v ...interface{}) {
	var logMsg string
	tagMsg := this.tagmsg(level)
	logMsg = fmt.Sprint(v...)
	logMsg = fmt.Sprintln(tagMsg, logMsg)
	fmt.Print(logMsg)
}

func (this CliLog) Outputf(level LogLevel, format string, v ...interface{}) {
	var logMsg string
	tagMsg := this.tagmsg(level)
	logMsg = fmt.Sprintf(format, v...)
	logMsg = fmt.Sprintln(tagMsg, logMsg)
	fmt.Print(logMsg)
}

func (this CliLog) tagmsg(level LogLevel) string {
	tagMsg := ""
	timeNow := ""
	if config.LOGER_PRINT_TIME {
		timeNow = fmt.Sprintf(" %s ", time.Now().Format("2006-01-02 15:04:05"))
	}
	switch level {
	case LL_DEBUG:
		tagMsg = "[DEBUG]" + timeNow
	case LL_INFO:
		tagMsg = "[INFO]" + timeNow
	case LL_WARNING:
		tagMsg = "[WARN]" + timeNow
	case LL_ERROR:
		tagMsg = "[ERR]" + timeNow
	case LL_CRITICAL:
		tagMsg = "[CRITICAL]" + timeNow
	}
	return tagMsg
}
