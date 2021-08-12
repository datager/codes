package log

import (
	"github.com/golang/glog"
)

type GLogLogger struct {
}

func NewGLogLogger() *GLogLogger {
	return &GLogLogger{}
}

func (logger *GLogLogger) Debug(args ...interface{}) {
	glog.V(4).Infoln(args...)
}

func (logger *GLogLogger) Debugln(args ...interface{}) {
	glog.V(4).Infoln(args...)
}

func (logger *GLogLogger) Debugf(format string, args ...interface{}) {
	glog.V(4).Infof(format, args...)
}

func (logger *GLogLogger) Info(args ...interface{}) {
	glog.Info(args...)
}

func (logger *GLogLogger) Infoln(args ...interface{}) {
	glog.Infoln(args...)
}

func (logger *GLogLogger) Infof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func (logger *GLogLogger) Warn(args ...interface{}) {
	glog.Warning(args...)
}

func (logger *GLogLogger) Warnln(args ...interface{}) {
	glog.Warning(args...)
}

func (logger *GLogLogger) Warnf(format string, args ...interface{}) {
	glog.Warningf(format, args...)
}

func (logger *GLogLogger) Error(args ...interface{}) {
	glog.Error(args...)
}

func (logger *GLogLogger) Errorln(args ...interface{}) {
	glog.Errorln(args...)
}

func (logger *GLogLogger) Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}
