package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

func getZapLog() *zap.Logger { return zapLog }

func LoggerInit() {

	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""

	zapLog, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, field ...zap.Field)  { getZapLog().Info(message, field...) }
func Debug(message string, field ...zap.Field) { getZapLog().Debug(message, field...) }
func Error(message string, field ...zap.Field) { getZapLog().Error(message, field...) }
func Warn(message string, field ...zap.Field)  { getZapLog().Warn(message, field...) }
func Fatal(message string, field ...zap.Field) { getZapLog().Fatal(message, field...) }
