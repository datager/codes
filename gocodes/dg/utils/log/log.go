package log

import "codes/gocodes/dg/utils"

var (
	DefaultLogger Logger
)

func init() {
	if utils.IsDevEnv() {
		logger, err := NewZapLogger()
		if err != nil {
			panic(err)
		}
		DefaultLogger = logger
	} else {
		DefaultLogger = NewGLogLogger()
	}
}

func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

func Debugln(args ...interface{}) {
	DefaultLogger.Debugln(args...)
}

func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

func Infoln(args ...interface{}) {
	DefaultLogger.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

func Warnln(args ...interface{}) {
	DefaultLogger.Warnln(args...)
}

func Warnf(format string, args ...interface{}) {
	DefaultLogger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	DefaultLogger.Errorln(args...)
}

func Errorln(args ...interface{}) {
	DefaultLogger.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	DefaultLogger.Errorf(format, args...)
}
