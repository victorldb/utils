package callbackmsg

//CallbackMsg --
type CallbackMsg interface {
	RegError(format string, v ...interface{})
	RegInfo(format string, v ...interface{})
}

// CallLogger --
type CallLogger interface {
	Error(format string, v ...interface{})
	Info(format string, v ...interface{})
}

//CallBackLog --
type CallBackLog struct {
	lf CallLogger
}

//NewCallBackLog --
func NewCallBackLog(lf CallLogger) *CallBackLog {
	newRM := &CallBackLog{lf: lf}
	return newRM
}

//RegError --
func (c *CallBackLog) RegError(format string, v ...interface{}) {
	c.lf.Error(format, v...)
}

//RegInfo --
func (c *CallBackLog) RegInfo(format string, v ...interface{}) {
	c.lf.Info(format, v...)
}
