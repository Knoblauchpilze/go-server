package rest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBodyFromResponseAs_InvalidResponse(t *testing.T) {
	assert := assert.New(t)

	var in foo
	err := GetBodyFromResponseAs(nil, &in)
	assert.Equal(err, ErrInvalidResponse)

	resp := http.Response{
		StatusCode: http.StatusBadRequest,
	}
	err = GetBodyFromResponseAs(&resp, &in)
	assert.Equal(err, ErrResponseIsError)
}

func TestGetBodyFromResponseAs_NoBody(t *testing.T) {
	assert := assert.New(t)

	resp := http.Response{
		StatusCode: http.StatusOK,
	}
	resp.Body = &mockBody{}

	var in foo
	err := GetBodyFromResponseAs(&resp, &in)
	assert.Equal(err, ErrFailedToGetBody)
}

func TestGetBodyFromResponseAs_InvalidBody(t *testing.T) {
	assert := assert.New(t)

	var in foo

	resp := generateResponseWithBody(nil)
	err := GetBodyFromResponseAs(resp, &in)
	assert.Equal(err, ErrBodyParsingFailed)

	resp = generateResponseWithBody([]byte("invalid"))
	err = GetBodyFromResponseAs(resp, &in)
	assert.Equal(err, ErrBodyParsingFailed)
}

func TestGetBodyFromResponseAs(t *testing.T) {
	assert := assert.New(t)

	in := foo{Bar: "bb", Baz: 12}
	data, _ := json.Marshal(in)
	resp := generateResponseWithBody(data)

	var out foo

	err := GetBodyFromResponseAs(resp, &out)
	assert.Nil(err)
	assert.Equal(out.Bar, in.Bar)
	assert.Equal(out.Baz, in.Baz)
}
