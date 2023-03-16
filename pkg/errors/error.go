package errors

import (
	"encoding/json"
	"fmt"
)

type ErrorWithCode interface {
	Code() ErrorCode
}

type errorImpl struct {
	hasCode bool
	Value   ErrorCode `json:"Code"`
	Message string
	Cause   error `json:",omitempty"`
}

func New(message string) error {
	return errorImpl{
		hasCode: false,
		Message: message,
	}
}

func NewCode(code ErrorCode) error {
	e := errorImpl{
		hasCode: true,
		Value:   code,
	}

	if msg, ok := errorsCodeToMessage[code]; ok {
		e.Message = msg
	} else {
		e.Message = defaultErrorMessage
	}

	return e
}

func Newf(format string, args ...interface{}) error {
	return errorImpl{
		hasCode: false,
		Message: fmt.Sprintf(format, args...),
	}
}

func Wrap(err error, message string) error {
	return errorImpl{
		Message: message,
		Cause:   err,
	}
}

func WrapCode(err error, code ErrorCode) error {
	e := errorImpl{
		hasCode: true,
		Value:   code,
		Cause:   err,
	}

	if msg, ok := errorsCodeToMessage[code]; ok {
		e.Message = msg
	} else {
		e.Message = defaultErrorMessage
	}

	return e
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errorImpl{
		Message: fmt.Sprintf(format, args...),
		Cause:   err,
	}
}

func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	ie, ok := err.(errorImpl)
	if !ok {
		return nil
	}

	return ie.Cause
}

func (e errorImpl) Error() string {
	var out string

	if e.hasCode {
		out += fmt.Sprintf("(%d) ", e.Value)
	}

	out += e.Message

	if e.Cause != nil {
		out += fmt.Sprintf(" (cause: %v)", e.Cause.Error())
	}

	return out
}

func (e errorImpl) Code() ErrorCode {
	if e.hasCode {
		return e.Value
	}

	return errGenericErrorCode
}

func (e errorImpl) MarshalJSON() ([]byte, error) {
	if !e.hasCode {
		return json.Marshal(struct {
			Message string
			Cause   error `json:",omitempty"`
		}{
			Message: e.Message,
			Cause:   e.Cause,
		})
	}

	return json.Marshal(struct {
		Code    ErrorCode
		Message string
		Cause   error `json:",omitempty"`
	}{
		Code:    e.Value,
		Message: e.Message,
		Cause:   e.Cause,
	})
}
