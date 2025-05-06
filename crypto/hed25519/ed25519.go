package hed25519

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"github.com/hootuu/hyle/hlog"
	"go.uber.org/zap"
)

func Random() ([]byte, []byte, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		hlog.Err("hyle.hed25519.Random", zap.Error(err))
		return nil, nil, errors.New("gen ed25519 key fail" + err.Error())
	}
	return pub, pri, nil
}

func Sign(priKey []byte, message []byte) []byte {
	return ed25519.Sign(priKey, message)
}

func Verify(pubKey []byte, message []byte, signature []byte) bool {
	return ed25519.Verify(pubKey, message, signature)
}
