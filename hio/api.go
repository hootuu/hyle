package hio

import "math/rand/v2"

type Api struct {
	BizID string `json:"biz_id"`
	Rand  int64  `json:"rand"`
}

func NewApi(bizID string) *Api {
	return &Api{
		BizID: bizID,
		Rand:  rand.Int64(),
	}
}

type ApiToken struct {
	Token            string `json:"token"`
	Refresh          string `json:"refresh"`
	TokenTimestamp   int64  `json:"token_timestamp"`
	RefreshTimestamp int64  `json:"refresh_timestamp"`
}
