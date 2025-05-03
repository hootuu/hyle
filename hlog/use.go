package hlog

import (
	"go.uber.org/zap"
)

func GetLogger(key string) *zap.Logger {
	return getLogger(key)
}

func Logger() *zap.Logger {
	return gInfoLogger
}

func Error() *zap.Logger {
	return gErrorLogger
}

func Info(msg string, fields ...zap.Field) {
	gInfoLogger.Info(msg, fields...)
}

func Err(msg string, fields ...zap.Field) {
	gErrorLogger.Error(msg, fields...)
}

var gErrorLogger *zap.Logger
var gInfoLogger *zap.Logger
var gElapseLogger *zap.Logger

func init() {
	gErrorLogger = getLogger("errors")
	gInfoLogger = getLogger("commons")
	gElapseLogger = getLogger("elapse")
}
