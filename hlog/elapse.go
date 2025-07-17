package hlog

import (
	"context"
	"github.com/hootuu/hyle/hcfg"
	"go.uber.org/zap"
	"time"
)

type ElapseLevel int

const (
	_ ElapseLevel = iota
	ElapsePackage
	ElapseComponent
	ElapseFunction
	ElapseDetail
)

const (
	TraceIdKey = "_trace_id_"
)

var gElapseLevel = ElapseDetail

func init() {
	iLevel := hcfg.GetInt("hlog.elapse.level", int(ElapseDetail))
	gElapseLevel = ElapseLevel(iLevel)
}

func IsElapsePackage() bool {
	return gElapseLevel >= ElapsePackage
}
func IsElapseComponent() bool {
	return gElapseLevel >= ElapseComponent
}
func IsElapseFunction() bool {
	return gElapseLevel >= ElapseFunction
}
func IsElapseDetail() bool {
	return gElapseLevel >= ElapseDetail
}

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
