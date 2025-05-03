package hcast

import "github.com/spf13/cast"

func ToString(val interface{}) string {
	return cast.ToString(val)
}

func ToInt(val interface{}) int {
	return cast.ToInt(val)
}

func ToIntE(val interface{}) (int, error) {
	return cast.ToIntE(val)
}

func ToBool(val interface{}) bool {
	return cast.ToBool(val)
}

func ToBoolE(val interface{}) (bool, error) {
	return cast.ToBoolE(val)
}

func ToInt64(val interface{}) int64 {
	return cast.ToInt64(val)
}

func ToInt64E(val interface{}) (int64, error) {
	return cast.ToInt64E(val)
}

func ToUint64(val interface{}) uint64 {
	return cast.ToUint64(val)
}

func ToUint64E(val interface{}) (uint64, error) {
	return cast.ToUint64E(val)
}
