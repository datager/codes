package log

type Logger interface {
	Debug(v ...interface{})
	Debugln(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infoln(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnln(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorln(v ...interface{})
	Errorf(format string, v ...interface{})
}
