package hlog

import (
	"context"
	"go.uber.org/zap"
)

type ElapseLogger struct {
	ctx    context.Context
	method string
	start  []zap.Field
	endFn  func() []zap.Field
}

func EL(ctx context.Context, method string) *ElapseLogger {
	return &ElapseLogger{
		ctx:    ctx,
		method: method,
	}
}

func (e *ElapseLogger) With(fields ...zap.Field) *ElapseLogger {
	e.start = append(e.start, fields...)
	return e
}

func (e *ElapseLogger) EndWith(fn func() []zap.Field) func() {
	if !IsElapseComponent() {
		return func() {}
	}

	return ElapseWithCtx(e.ctx, e.method, F(e.start...), func() []zap.Field {
		fields := fn()
		return filterEmptyFields(fields)
	})
}

func filterEmptyFields(fields []zap.Field) []zap.Field {
	var result []zap.Field
	for _, field := range fields {
		if field.Interface == nil {
			continue
		}
		result = append(result, field)
	}
	return result
}
