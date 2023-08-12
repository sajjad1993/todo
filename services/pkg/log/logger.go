package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(messages ...interface{})
	Info(messages ...interface{})
	Warning(messages ...interface{})
	Error(messages ...interface{})
	Fatal(messages ...interface{})
	Panic(messages ...interface{})
	Infof(template string, args ...interface{})
}

func NewCustomLogger(cores ...zapcore.Core) Logger {
	core := zapcore.NewTee(cores...)
	levelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= ErrorLevel
	})
	logger := zap.New(core, zap.AddStacktrace(levelEnabler))
	zapSuger := logger.Sugar()
	return &zapLogger{logger: zapSuger}
}

func NewLogger() Logger {

	return NewCustomLogger(
		NewStdoutCore(),
	)
}
