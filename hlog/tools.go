package hlog

import "go.uber.org/zap"

func nullZ() []zap.Field {
	return nil
}

func F(arr ...zap.Field) func() []zap.Field {
	if len(arr) == 0 {
		return nullZ
	}

	return func() []zap.Field {
		return arr
	}
}

func E(err error, arr ...zap.Field) func() []zap.Field {
	var fields []zap.Field
	if err == nil {
		fields = append(fields, zap.Error(err))
	}
	if len(arr) > 0 {
		fields = append(fields, arr...)
	}
	if len(fields) > 0 {
		return func() []zap.Field {
			return fields
		}
	}
	return nullZ
}
