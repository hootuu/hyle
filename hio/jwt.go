package hio

type JwtToken struct {
	Token             string `json:"token"`
	Refresh           string `json:"refresh"`
	TokenExpiration   int64  `json:"token_exp"`
	RefreshExpiration int64  `json:"refresh_exp"`
}
