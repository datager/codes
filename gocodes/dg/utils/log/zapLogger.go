package log

import (
	"codes/gocodes/dg/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	*zap.SugaredLogger
}

func NewZapLogger() (*ZapLogger, error) {
	var logger *zap.Logger
	var err error
	if utils.IsDevEnv() {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = config.Build()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ZapLogger{logger.Sugar()}, nil
}

func (logger *ZapLogger) Debugln(args ...interface{}) {
	logger.Debug(args...)
}

func (logger *ZapLogger) Infoln(args ...interface{}) {
	logger.Info(args...)
}

func (logger *ZapLogger) Warnln(args ...interface{}) {
	logger.Warn(args...)
}

func (logger *ZapLogger) Errorln(args ...interface{}) {
	logger.Error(args...)
}
