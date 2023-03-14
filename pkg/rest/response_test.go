package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func generateResponseWithBody(body interface{}) *http.Response {
	resp := http.Response{
		StatusCode: http.StatusOK,
	}

	in := NewSuccessResponse(uuid.UUID{})
	if body != nil {
		in.WithDetails(body)
	}

	data, _ := json.Marshal(in)

	rdr := bytes.NewReader(data)
	resp.Body = io.NopCloser(rdr)

	return &resp
}

func TestGetBodyFromHttpResponseAs_InvalidResponse(t *testing.T) {
	assert := assert.New(t)

	var in foo
	err := GetBodyFromHttpResponseAs(nil, &in)
	assert.Equal(err, ErrNoResponse)

	resp := http.Response{
		StatusCode: http.StatusBadRequest,
	}
	err = GetBodyFromHttpResponseAs(&resp, &in)
	assert.Equal(err, ErrResponseIsError)

	resp.StatusCode = http.StatusOK
	rdr := bytes.NewReader([]byte("haha"))
	resp.Body = io.NopCloser(rdr)

	err = GetBodyFromHttpResponseAs(&resp, &in)
	assert.Equal(err, ErrInvalidResponse)
}

func TestGetBodyFromHttpResponseAs_NoBody(t *testing.T) {
	assert := assert.New(t)

	resp := http.Response{
		StatusCode: http.StatusOK,
	}
	resp.Body = &mockBody{}

	var in foo
	err := GetBodyFromHttpResponseAs(&resp, &in)
	assert.Equal(err, ErrFailedToGetBody)
}

func TestGetBodyFromHttpResponseAs_InvalidBody(t *testing.T) {
	assert := assert.New(t)

	var in foo

	resp := generateResponseWithBody(nil)
	err := GetBodyFromHttpResponseAs(resp, &in)
	assert.Equal(err, ErrBodyParsingFailed)

	resp = generateResponseWithBody("invalid")
	err = GetBodyFromHttpResponseAs(resp, &in)
	assert.Equal(err, ErrBodyParsingFailed)
}

func TestGetBodyFromHttpResponseAs(t *testing.T) {
	assert := assert.New(t)

	in := foo{Bar: "bb", Baz: 12}
	resp := generateResponseWithBody(in)

	var out foo

	err := GetBodyFromHttpResponseAs(resp, &out)
	assert.Nil(err)
	assert.Equal(out.Bar, in.Bar)
	assert.Equal(out.Baz, in.Baz)
}
