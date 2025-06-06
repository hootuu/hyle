package collar

import (
	"github.com/hootuu/hyle/hlog"
	"go.uber.org/zap"
	"strings"
)

const (
	split = ":"
)

type Collar string

func Build(code string, id string) Collar {
	return Collar(code + split + id)
}

func (c Collar) Parse() (string, string) {
	arr := strings.Split(string(c), split)
	if len(arr) != 2 {
		hlog.Err("hypes.collar.Parse: invalid Collar", zap.String("collar", string(c)))
		return "", ""
	}
	return arr[0], arr[1]
}
