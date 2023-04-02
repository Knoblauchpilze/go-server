package connection

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
)

func NewPostRequest(url string, headers http.Header, contentType string, body interface{}) RequestWrapper {
	builder := func(ri *requestImpl) (*http.Request, error) {
		return buildPostRequest(ri, contentType, body)
	}

	return newRequest(url, headers, builder)
}

func buildPostRequest(ri *requestImpl, contentType string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest("POST", ri.url, nil)
	if err != nil {
		return req, errors.WrapCode(err, errors.ErrGetRequestFailed)
	}

	data, err := json.Marshal(body)
	if err != nil {
		return req, errors.WrapCode(err, errors.ErrPostInvalidData)
	}

	req.Header = ri.headers
	req.Header.Set("Content-Type", contentType)
	req.Body = io.NopCloser(bytes.NewReader(data))

	return req, nil
}
