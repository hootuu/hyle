package collar

import (
	"fmt"
	"github.com/mr-tron/base58"
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
	arr := strings.Split(string(src), split)
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
	arr := strings.Split(string(src), split)
	if len(arr) != 2 {
		return "", fmt.Errorf("invalid collar: %s", src)
	}
	return Build(arr[0], arr[1]), nil
}

func (k Link) Display() string {
	code, id, _ := k.ToCodeID()
	return fmt.Sprintf("%s:%s", code, id)
}
