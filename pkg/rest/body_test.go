package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type foo struct {
	Bar string
	Baz int
}

var errSomeError = fmt.Errorf("some error")
var errAnotherError = fmt.Errorf("another error")

type mockBody struct{}

func (mb mockBody) Read(p []byte) (n int, err error) {
	return 0, errSomeError
}

func (mb mockBody) Close() error {
	return errAnotherError
}

func generateRequestWithBody(body []byte) *http.Request {
	req := http.Request{}

	rdr := bytes.NewReader(body)
	req.Body = io.NopCloser(rdr)

	return &req
}

func TestGetBodyAsFromRequest_NoBody(t *testing.T) {
	assert := assert.New(t)

	req := http.Request{}
	req.Body = &mockBody{}

	var in foo
	err := GetBodyFromRequestAs(&req, &in)
	assert.Equal(err, ErrFailedToGetBody)
}

func TestGetBodyAsFromRequest_InvalidBody(t *testing.T) {
	assert := assert.New(t)

	var in foo

	req := generateRequestWithBody(nil)
	err := GetBodyFromRequestAs(req, &in)
	assert.Equal(err, ErrBodyParsingFailed)

	req = generateRequestWithBody([]byte("invalid"))
	err = GetBodyFromRequestAs(req, &in)
	assert.Equal(err, ErrBodyParsingFailed)
}

func TestGetBodyAsFromRequest(t *testing.T) {
	assert := assert.New(t)

	in := foo{Bar: "bb", Baz: 12}
	data, _ := json.Marshal(in)
	req := generateRequestWithBody(data)

	var out foo

	err := GetBodyFromRequestAs(req, &out)
	assert.Nil(err)
	assert.Equal(out.Bar, in.Bar)
	assert.Equal(out.Baz, in.Baz)
}

func GetBodyAsFromRequest(req *http.Request, out interface{}) error {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return ErrFailedToGetBody
	}

	err = json.Unmarshal(data, out)
	if err != nil {
		return err
	}

	return nil
}
