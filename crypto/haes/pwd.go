package haes

import (
	"crypto/sha256"
	"github.com/hootuu/hyle/hlog"
	"go.uber.org/zap"
)

type Password []byte

func (pwd Password) Cover(src []byte) ([]byte, error) {
	hash := sha256.Sum256(pwd)
	bytes, err := Encrypt(src, hash[:])
	if err != nil {
		hlog.Err("haes.pwd.Cover", zap.Error(err))
		return nil, err
	}
	return bytes, nil
}
func (pwd Password) Uncover(src []byte) ([]byte, error) {
	hash := sha256.Sum256(pwd)
	bytes, err := Decrypt(src, hash[:])
	if err != nil {
		hlog.Err("haes.pwd.Uncover", zap.Error(err))
		return nil, err
	}
	return bytes, nil
}
