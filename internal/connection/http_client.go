package connection

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
)

func performGetRequest(url string, headers map[string][]string) (*http.Response, error) {
	var resp *http.Response
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resp, errors.WrapCode(err, errors.ErrGetRequestFailed)
	}

	req.Header = headers

	resp, err = client.Do(req)
	if err != nil {
		return resp, errors.WrapCode(err, errors.ErrGetRequestFailed)
	}

	return resp, nil
}

func performPostRequest(url string, headers map[string][]string, contentType string, body interface{}) (*http.Response, error) {
	var resp *http.Response
	client := http.Client{}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return resp, errors.WrapCode(err, errors.ErrPostRequestFailed)
	}

	data, err := json.Marshal(body)
	if err != nil {
		return resp, errors.WrapCode(err, errors.ErrPostInvalidData)
	}

	req.Header = headers
	req.Header.Set("Content-Type", contentType)
	req.Body = io.NopCloser(bytes.NewReader(data))

	resp, err = client.Do(req)
	if err != nil {
		return resp, errors.WrapCode(err, errors.ErrPostRequestFailed)
	}

	return resp, nil
}
