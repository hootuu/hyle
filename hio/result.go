package hio

const (
	ResultOK = 1
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"-"`
}

type Result[T any] struct {
	Success bool   `json:"success"`
	Data    *T     `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

func OK[T any](data *T) *Result[T] {
	return &Result[T]{
		Success: true,
		Data:    data,
	}
}

func NG[T any](code int, message string, err error) *Result[T] {
	return &Result[T]{
		Success: false,
		Error:   &Error{Code: code, Message: message, Error: err},
	}
}
