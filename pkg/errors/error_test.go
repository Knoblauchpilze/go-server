package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errSomeError = fmt.Errorf("some error")

func TestError_New(t *testing.T) {
	assert := assert.New(t)

	err := New("haha")

	impl, ok := err.(errorImpl)
	assert.True(ok)
	assert.Equal(impl.Message, "haha")
	assert.Nil(impl.Cause)
}

func TestError_Newf(t *testing.T) {
	assert := assert.New(t)

	err := Newf("haha %d", 22)

	impl, ok := err.(errorImpl)
	assert.True(ok)
	assert.Equal(impl.Message, "haha 22")
	assert.Nil(impl.Cause)
}

func TestError_Wrap(t *testing.T) {
	assert := assert.New(t)

	err := Wrap(errSomeError, "context")

	impl, ok := err.(errorImpl)
	assert.True(ok)
	assert.Equal(impl.Message, "context")
	assert.Equal(impl.Cause, errSomeError)
}

func TestError_Wrapf(t *testing.T) {
	assert := assert.New(t)

	err := Wrapf(errSomeError, "context %d", -44)

	impl, ok := err.(errorImpl)
	assert.True(ok)
	assert.Equal(impl.Message, "context -44")
	assert.Equal(impl.Cause, errSomeError)
}

func TestError_Unwrap(t *testing.T) {
	assert := assert.New(t)

	err := Unwrap(errSomeError)
	assert.Nil(err)

	err = New("haha")
	cause := Unwrap(err)
	assert.Nil(cause)

	err = Wrap(errSomeError, "haha")
	cause = Unwrap(err)
	assert.Equal(cause, errSomeError)

	causeOfCause := Unwrap(cause)
	assert.Nil(causeOfCause)
}

func TestError_Error(t *testing.T) {
	assert := assert.New(t)

	err := Wrapf(errSomeError, "context %d", -44)

	expected := "context -44 (cause: some error)"
	assert.Equal(err.Error(), expected)
}
