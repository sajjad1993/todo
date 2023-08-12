package log

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func (zapLog *zapLogger) Debug(messages ...interface{}) {
	zapLog.logger.Debug(messages)
}

func (zapLog *zapLogger) Info(messages ...interface{}) {
	zapLog.logger.Info(messages)
}

func (z zapLogger) Infof(template string, args ...interface{}) {
	z.logger.Infof(template, args...)
}

func (zapLog *zapLogger) Warning(messages ...interface{}) {
	zapLog.logger.Warn(messages)
}

func (zapLog *zapLogger) Error(messages ...interface{}) {
	zapLog.logger.Error(messages)
}

func (zapLog *zapLogger) Fatal(messages ...interface{}) {
	zapLog.logger.Fatal(messages)
}

func (zapLog *zapLogger) Panic(messages ...interface{}) {
	zapLog.logger.Panic(messages)
}
