package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// standard output.
func NewStdoutCore() zapcore.Core {
	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	syncWriter := zapcore.AddSync(os.Stdout)
	levelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})

	return zapcore.NewCore(jsonEncoder, syncWriter, levelEnabler)
}
