package rest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBodyFromHttpRequestAs_NoBody(t *testing.T) {
	assert := assert.New(t)

	req := http.Request{}
	req.Body = &mockBody{}

	var in foo
	err := GetBodyFromHttpRequestAs(&req, &in)
	assert.Equal(err, ErrFailedToGetBody)
}

func TestGetBodyFromHttpRequestAs_InvalidBody(t *testing.T) {
	assert := assert.New(t)

	var in foo

	req := generateRequestWithBody(nil)
	err := GetBodyFromHttpRequestAs(req, &in)
	assert.Equal(err, ErrBodyParsingFailed)

	req = generateRequestWithBody([]byte("invalid"))
	err = GetBodyFromHttpRequestAs(req, &in)
	assert.Equal(err, ErrBodyParsingFailed)
}

func TestGetBodyFromHttpRequestAs(t *testing.T) {
	assert := assert.New(t)

	in := foo{Bar: "bb", Baz: 12}
	data, _ := json.Marshal(in)
	req := generateRequestWithBody(data)

	var out foo

	err := GetBodyFromHttpRequestAs(req, &out)
	assert.Nil(err)
	assert.Equal(out.Bar, in.Bar)
	assert.Equal(out.Baz, in.Baz)
}
