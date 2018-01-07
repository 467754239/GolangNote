package log

type LogLevel int

type Log interface {
	Output(level LogLevel, v ...interface{})
	Outputf(level LogLevel, format string, v ...interface{})
}

const (
	LL_DEBUG LogLevel = iota
	LL_INFO
	LL_WARNING
	LL_ERROR
	LL_CRITICAL
)
