package herr

import (
	"errors"
	"fmt"
)

type Code uint32

type Error struct {
	Code    Code   `bson:"code"  json:"code"`
	Message string `bson:"message" json:"message"`

	Err error `bson:"-" json:"-"` // Native error
}

func Of(code Code, message string, nativeErr ...error) *Error {
	err := &Error{
		Code:    code,
		Message: message,
	}
	if len(nativeErr) > 0 && nativeErr[0] != nil {
		err.Err = nativeErr[0]
	}
	return err
}

func (e *Error) Error() string {
	if e.Err != nil && !errors.Is(e, e.Err) {
		return fmt.Sprintf("[%d] %s. [ %s ]", e.Code, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func Assert(assert bool, code Code, msg string) *Error {
	if assert {
		return nil
	}
	return Of(code, msg)
}

func (e *Error) GetCode() Code {
	return e.Code
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) Native() error {
	if e.Err == nil {
		return errors.New(fmt.Sprintf("[%d] %s", e.Code, e.Message))
	}
	return e.Err
}
