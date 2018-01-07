package log

type DisabledLog struct{}

func (this DisabledLog) Output(level LogLevel, v ...interface{}) {}

func (this DisabledLog) Outputf(level LogLevel, format string, v ...interface{}) {}
