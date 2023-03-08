package rest

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateRequestWithHeader() http.Request {
	return http.Request{
		Header: make(http.Header),
	}
}

func TestGetHeaderFromRequest_NoHeader(t *testing.T) {
	assert := assert.New(t)

	req := http.Request{}

	_, err := GetHeaderFromRequest(&req, "foo")
	assert.Equal(err, ErrNoSuchHeader)
}

func TestGetHeaderFromRequest_OneValue(t *testing.T) {
	assert := assert.New(t)

	headerValues := []string{"haha"}

	req := generateRequestWithHeader()
	req.Header["foo"] = headerValues

	out, err := GetHeaderFromRequest(&req, "foo")
	assert.Nil(err)
	assert.Equal(len(headerValues), len(out))

	for id, expectedHeader := range headerValues {
		assert.Equal(expectedHeader, out[id])
	}
}

func TestGetHeaderFromRequest_AnotherValue(t *testing.T) {
	assert := assert.New(t)

	req := generateRequestWithHeader()
	req.Header["food"] = []string{"haha"}

	_, err := GetHeaderFromRequest(&req, "foo")
	assert.Equal(err, ErrNoSuchHeader)
}

func TestGetHeaderFromRequest_TwoValues(t *testing.T) {
	assert := assert.New(t)

	fooHeaderValues := []string{"haha"}
	barHeaderValues := []string{"hihi"}

	req := generateRequestWithHeader()
	req.Header["foo"] = fooHeaderValues
	req.Header["bar"] = barHeaderValues

	out, err := GetHeaderFromRequest(&req, "foo")
	assert.Nil(err)
	assert.Equal(len(fooHeaderValues), len(out))

	for id, expectedHeader := range fooHeaderValues {
		assert.Equal(expectedHeader, out[id])
	}
}

func TestGetSingleHeaderFromRequest_NoHeader(t *testing.T) {
	assert := assert.New(t)

	req := http.Request{}

	_, err := GetSingleHeaderFromRequest(&req, "foo")
	assert.Equal(err, ErrNoSuchHeader)
}

func TestGetSingleHeaderFromRequest_OneValue(t *testing.T) {
	assert := assert.New(t)

	headerValues := []string{"haha"}

	req := generateRequestWithHeader()
	req.Header["foo"] = headerValues

	out, err := GetSingleHeaderFromRequest(&req, "foo")
	assert.Nil(err)
	assert.Equal("haha", out)
}

func TestGetSingleHeaderFromRequest_MultipleValues(t *testing.T) {
	assert := assert.New(t)

	headerValues := []string{"haha", "hihi"}

	req := generateRequestWithHeader()
	req.Header["foo"] = headerValues

	_, err := GetSingleHeaderFromRequest(&req, "foo")
	assert.Equal(err, ErrNonUniqueHeader)
}
