package hio

import (
	"encoding/json"
	"github.com/hootuu/hyle/herr"
	"math/rand/v2"
	"time"
)

const (
	RespMarshalErr   herr.Code = 10101022
	RespUnmarshalErr herr.Code = 10101023
)

type Response[T any] struct {
	ReqID     string      `json:"req_id"`
	Success   bool        `json:"success"`
	Data      *T          `json:"data,omitempty"`
	Error     *herr.Error `json:"error,omitempty"`
	Timestamp int64       `json:"timestamp"`
	Nonce     int64       `json:"nonce"`
}

func NewResponse[T any](reqID string, data *T) *Response[T] {
	randInt64 := rand.Int64()
	return &Response[T]{
		ReqID:     reqID,
		Success:   true,
		Data:      data,
		Error:     nil,
		Timestamp: time.Now().UnixMilli(),
		Nonce:     randInt64,
	}
}

func FailResponse[T any](reqID string, err *herr.Error) *Response[T] {
	randInt64 := rand.Int64()
	return &Response[T]{
		ReqID:     reqID,
		Success:   false,
		Data:      nil,
		Error:     err,
		Timestamp: time.Now().UnixMilli(),
		Nonce:     randInt64,
	}
}

func (r *Response[T]) Marshal() ([]byte, *herr.Error) {
	byteArr, err := json.Marshal(r)
	if err != nil {
		return nil, herr.Of(RespMarshalErr, "marshal response failed: "+err.Error())
	}
	return byteArr, nil
}

func (r *Response[T]) Unmarshal(data []byte) *herr.Error {
	if err := json.Unmarshal(data, r); err != nil {
		return herr.Of(RespUnmarshalErr, "unmarshal response failed: "+err.Error())
	}
	return nil
}

func ResponseUnmarshal[T any](data []byte) (*Response[T], *herr.Error) {
	var resp Response[T]

	if err := resp.Unmarshal(data); err != nil {
		return nil, err
	}
	return &resp, nil
}
