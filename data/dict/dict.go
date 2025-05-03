package dict

import "github.com/stretchr/objx"

type Dict = objx.Map

func New(data interface{}) Dict {
	return objx.New(data)
}
