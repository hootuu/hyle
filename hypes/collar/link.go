package collar

import (
	"fmt"
	"github.com/hootuu/hyle/hlog"
	"github.com/mr-tron/base58"
	"go.uber.org/zap"
	"strings"
)

type Link string

type LinkDict struct {
	Code string `json:"code"`
	ID   ID     `json:"id"`
}

func (k Link) Str() string {
	return string(k)
}

func (k Link) ToCodeID() (string, string, error) {
	src, err := base58.Decode(k.Str())
	if err != nil {
		return "", "", err
	}
	arr := strings.SplitN(string(src), split, 2)
	if len(arr) != 2 {
		return "", "", fmt.Errorf("invalid collar: %s", src)
	}
	return arr[0], arr[1], nil
}

func (k Link) MustToCodeID() (string, string) {
	code, id, err := k.ToCodeID()
	if err != nil {
		return "", ""
	}
	return code, id
}

func (k Link) MustToCode() string {
	code, _, err := k.ToCodeID()
	if err != nil {
		return ""
	}
	return code
}

func (k Link) MustToID() ID {
	_, id, err := k.ToCodeID()
	if err != nil {
		return ""
	}
	return id
}

func (k Link) MustToDict() *LinkDict {
	code, id := k.MustToCodeID()
	return &LinkDict{
		Code: code,
		ID:   id,
	}
}

func (k Link) ToCollar() (Collar, error) {
	src, err := base58.Decode(k.Str())
	if err != nil {
		return "", err
	}
	arr := strings.SplitN(string(src), split, 2)
	if len(arr) != 2 {
		return "", fmt.Errorf("invalid collar: %s", src)
	}
	return Build(arr[0], arr[1]), nil
}

func (k Link) Display() string {
	code, id, err := k.ToCodeID()
	if err != nil {
		hlog.Fix("invalid link:", zap.String("link", k.Str()), zap.Error(err))
	}
	return fmt.Sprintf("%s:%s", code, id)
}
