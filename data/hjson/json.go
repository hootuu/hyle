package hjson

import (
	"encoding/json"
	"github.com/hootuu/hyle/hlog"
	"go.uber.org/zap"
)

func ToBytes(obj interface{}) ([]byte, error) {
	if obj == nil {
		return []byte{}, nil
	}
	return json.Marshal(obj)
}

func MustToBytes(obj interface{}) []byte {
	bytes, err := ToBytes(obj)
	if err != nil {
		hlog.Err("hyle.json.MustToBytes", zap.Error(err))
		return []byte{}
	}
	return bytes
}

func ToString(obj interface{}) (string, error) {
	bData, err := ToBytes(obj)
	if err != nil {
		return "", err
	}
	return string(bData), nil
}

func MustToString(obj interface{}) string {
	str, err := ToString(obj)
	if err != nil {
		hlog.Err("hyle.json.MustToString", zap.Error(err))
		return ""
	}
	return str
}

func FromBytes[T any](bData []byte) (*T, error) {
	var obj T
	err := json.Unmarshal(bData, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func MustFromBytes[T any](bData []byte) *T {
	ptrObj, err := FromBytes[T](bData)
	if err != nil {
		hlog.Err("hyle.json.MustFromBytes",
			zap.Error(err),
			zap.String("json", string(bData)))
		return nil
	}
	if ptrObj == nil {
		var obj T
		return &obj
	}
	return ptrObj
}

func FromString[T any](str string) (*T, error) {
	return FromBytes[T]([]byte(str))
}

func MustFromString[T any](str string) *T {
	obj, err := FromString[T](str)
	if err != nil {
		return nil
	}
	return obj
}

func OfBytes[T any](ptr *T, bData []byte) error {
	return json.Unmarshal(bData, ptr)
}

func MustOfBytes[T any](ptr *T, bData []byte) {
	err := OfBytes[T](ptr, bData)
	if err != nil {
		hlog.Err("hyle.json.MustOfBytes", zap.Error(err))
	}
}

func OfString[T any](ptr *T, str string) error {
	return OfBytes[T](ptr, []byte(str))
}

func MustOfString[T any](ptr *T, str string) {
	err := OfString[T](ptr, str)
	if err != nil {
		hlog.Err("hyle.json.MustOfString", zap.Error(err))
	}
}
