package dict

import (
	"github.com/hootuu/hyle/data/hjson"
	"github.com/stretchr/objx"
)

type Dict = objx.Map

func New(data interface{}) Dict {
	if data == nil {
		return objx.New(make(map[string]interface{}))
	}
	return objx.New(data)
}

func NewDict() Dict {
	return objx.New(make(map[string]interface{}))
}

func NewFrom(obj any) Dict {
	if obj == nil {
		return NewDict()
	}
	bytes := hjson.MustToBytes(obj)
	m := hjson.MustFromBytes[map[string]any](bytes)
	return New(*m)
}
