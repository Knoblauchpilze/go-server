package errors

import "fmt"

type errorImpl struct {
	Message string
	Cause   error
}

func New(message string) error {
	return errorImpl{
		Message: message,
	}
}

func Newf(format string, args ...interface{}) error {
	return errorImpl{
		Message: fmt.Sprintf(format, args...),
	}
}

func Wrap(err error, message string) error {
	return errorImpl{
		Message: message,
		Cause:   err,
	}
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
	out := e.Message
	if e.Cause != nil {
		out += fmt.Sprintf(" (cause: %v)", e.Cause.Error())
	}

	return out
}
