package hlog

import (
	"context"
	"go.uber.org/zap"
	"time"
)

const (
	TraceIdKey = "_trace_id_"
)

func Elapse(fun string, fix ...func() []zap.Field) func() {
	start := time.Now()
	var prefixFields []zap.Field
	if len(fix) > 0 {
		prefixFields = fix[0]()
	}
	if len(prefixFields) > 0 {
		gElapseLogger.Info("==>"+fun, prefixFields...)
	} else {
		gElapseLogger.Info("==>" + fun)
	}

	return func() {
		elapse := time.Since(start)
		suffixFields := []zap.Field{zap.Int64("_elapse", elapse.Milliseconds())}
		if len(fix) > 1 {
			fs := fix[1]()
			if len(fs) > 1 {
				suffixFields = append(suffixFields, fs...)
			}
		}
		if len(suffixFields) > 0 {
			gElapseLogger.Info("<=="+fun, suffixFields...)
		} else {
			gElapseLogger.Info("<==" + fun)
		}
	}
}

func ElapseWithCtx(ctx context.Context, fun string, fix ...func() []zap.Field) func() {
	start := time.Now()
	var prefixFields []zap.Field
	if len(fix) > 0 {
		prefixFields = fix[0]()
	}
	if traceIdObj := ctx.Value(TraceIdKey); traceIdObj != nil {
		prefixFields = append(prefixFields, zap.String(TraceIdKey, traceIdObj.(string)))
	}
	if len(prefixFields) > 0 {
		gElapseLogger.Info(">"+fun, prefixFields...)
	} else {
		gElapseLogger.Info(">" + fun)
	}

	return func() {
		elapse := time.Since(start)
		suffixFields := []zap.Field{zap.Int64("_elapse_", elapse.Milliseconds())}
		if len(fix) > 1 {
			fs := fix[1]()
			if len(fs) > 1 {
				suffixFields = append(suffixFields, fs...)
			}
		}
		if traceIdObj := ctx.Value(TraceIdKey); traceIdObj != nil {
			suffixFields = append(suffixFields, zap.String(TraceIdKey, traceIdObj.(string)))
		}
		if len(suffixFields) > 0 {
			gElapseLogger.Info("<"+fun, suffixFields...)
		} else {
			gElapseLogger.Info("<" + fun)
		}
	}
}
