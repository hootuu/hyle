package hlog

import (
	"context"
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

func TraceErr(msg string, ctx context.Context, err error, fields ...zap.Field) {
	fullFields := []zap.Field{
		TraceInfo(ctx), zap.Error(err),
	}
	if len(fields) > 0 {
		fullFields = append(fullFields, fields...)
	}
	gElapseLogger.Error(msg, fullFields...)
}

func Fix(msg string, fields ...zap.Field) {
	gFixLogger.Error(msg, fields...)
	Err(msg, fields...)
}

func TraceFix(msg string, ctx context.Context, err error, fields ...zap.Field) {
	fullFields := []zap.Field{
		TraceInfo(ctx), zap.Error(err),
	}
	if len(fields) > 0 {
		fullFields = append(fullFields, fields...)
	}
	gFixLogger.Error(msg, fullFields...)
	Err(msg, fields...)
}

var gErrorLogger *zap.Logger
var gInfoLogger *zap.Logger
var gElapseLogger *zap.Logger
var gFixLogger *zap.Logger

func init() {
	gErrorLogger = getLogger("errors")
	gInfoLogger = getLogger("commons")
	gElapseLogger = getLogger("elapse")
	gFixLogger = getLogger("fix")
}
