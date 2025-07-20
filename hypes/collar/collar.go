package collar

import (
	"fmt"
	"github.com/hootuu/hyle/crypto/hmd5"
	"github.com/hootuu/hyle/hlog"
	"github.com/mr-tron/base58"
	"go.uber.org/zap"
	"strings"
)

const (
	split = ":"
)

type Collar string
type ID = string
type SafeID string

func Build(code string, id string) Collar {
	str := code + split + id
	if len(str) > 64 {
		hlog.Err("collar.length > 64 chars ", zap.String("code", code), zap.String("id", id))
		str = str[:64]
	}
	return Collar(str)
}

func FromID(id ID) (Collar, error) {
	src, err := base58.Decode(id)
	if err != nil {
		return "", err
	}
	arr := strings.Split(string(src), split)
	if len(arr) != 2 {
		return "", fmt.Errorf("invalid collar: %s", src)
	}
	return Build(arr[0], arr[1]), nil
}

func MustFromID(id ID) Collar {
	c, err := FromID(id)
	if err != nil {
		hlog.Fix("collar.mustFromID", zap.String("id", string(id)), zap.Error(err))
	}
	return c
}

func MustParse(id ID, call func(code string, id string)) {
	col, err := FromID(id)
	if err != nil {
		hlog.Fix("collar.mustParse", zap.String("id", id), zap.Error(err))
		return
	}
	code, id := col.Parse()
	call(code, id)
}

func MultiMustParse(call func(code string, id string), ids ...ID) {
	if len(ids) == 0 {
		return
	}
	for _, idStr := range ids {
		col, err := FromID(idStr)
		if err != nil {
			hlog.Fix("collar.mustParse", zap.String("id", idStr), zap.Error(err))
			continue
		}
		code, id := col.Parse()
		call(code, id)
	}
}

func (c Collar) Parse() (string, string) {
	arr := strings.Split(string(c), split)
	if len(arr) != 2 {
		hlog.Err("hypes.collar.Parse: invalid Collar", zap.String("collar", string(c)))
		return "", ""
	}
	return arr[0], arr[1]
}

func (c Collar) ToString() string {
	return string(c)
}

func (c Collar) ToID() ID {
	return base58.Encode([]byte(c))
}

func (c Collar) ToSafeID() ID {
	return hmd5.MD5(hmd5.MD5(c.ToID()))
}

func (c Collar) Link() Link {
	return Link(base58.Encode([]byte(c)))
}
