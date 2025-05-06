package hio

import (
	"encoding/hex"
	"encoding/json"
	"github.com/hootuu/hyle/crypto/hed25519"
	"github.com/hootuu/hyle/data/idx"
	"github.com/hootuu/hyle/herr"
	"math/rand/v2"
	"time"
)

const (
	ReqSerializeErr herr.Code = 10101010
	ReqVerifyErr    herr.Code = 10101011
	ReqMarshalErr   herr.Code = 10101012
	ReqUnmarshalErr herr.Code = 10101013
)

type Request[T any] struct {
	ReqID     string `json:"req_id"`
	TokenID   string `json:"token_id"`
	Timestamp int64  `json:"timestamp"`
	Nonce     int64  `json:"nonce"`
	Signature string `json:"signature"`
	Data      *T     `json:"data,omitempty"`
}

func NewRequest[T any](tokenID string, data *T) *Request[T] {
	randInt64 := rand.Int64()
	return &Request[T]{
		Data:      data,
		ReqID:     idx.New(),
		TokenID:   tokenID,
		Timestamp: time.Now().UnixMilli(),
		Nonce:     randInt64,
	}
}

func (req *Request[T]) Sign(priKey []byte) *herr.Error {
	bytes, err := req.doSerialize()
	if err != nil {
		return err
	}
	sig := hed25519.Sign(priKey, bytes)
	req.Signature = hex.EncodeToString(sig)
	return nil
}

func (req *Request[T]) Verify(pubKey []byte) *herr.Error {
	bytes, err := req.doSerialize()
	if err != nil {
		return err
	}
	bytesSign, nErr := hex.DecodeString(req.Signature)
	if nErr != nil {
		return err
	}
	valid := hed25519.Verify(pubKey, bytes, bytesSign)
	if !valid {
		return herr.Of(ReqVerifyErr, "invalid signature")
	}
	return nil
}

func (req *Request[T]) Marshal() ([]byte, *herr.Error) {
	byteArr, err := json.Marshal(req)
	if err != nil {
		return nil, herr.Of(ReqMarshalErr, "marshal request failed", err)
	}
	return byteArr, nil
}

func (req *Request[T]) Unmarshal(data []byte) *herr.Error {
	if err := json.Unmarshal(data, req); err != nil {
		return herr.Of(ReqUnmarshalErr, "unmarshal request failed", err)
	}
	return nil
}

func RequestUnmarshal[T any](data []byte) (*Request[T], *herr.Error) {
	var req Request[T]
	if err := req.Unmarshal(data); err != nil {
		return nil, err
	}
	return &req, nil
}

func (req *Request[T]) doSerialize() ([]byte, *herr.Error) {
	serializeStr, err := OrderedSerialize(map[string]interface{}{
		"req_id":    req.ReqID,
		"token_id":  req.TokenID,
		"data":      req.Data,
		"timestamp": req.Timestamp,
		"nonce":     req.Nonce,
	})
	if err != nil {
		return nil, herr.Of(ReqSerializeErr, "request serialize failed:"+err.Error())
	}
	return []byte(serializeStr), nil
}
