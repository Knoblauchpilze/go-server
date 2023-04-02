package connection

import (
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
)

type requestBuilder func(ri *requestImpl) (*http.Request, error)

type requestImpl struct {
	url     string
	headers http.Header
	builder requestBuilder
}

func newRequest(url string, headers http.Header, builder requestBuilder) RequestWrapper {
	return &requestImpl{
		url:     url,
		headers: headers,
		builder: builder,
	}
}

func (ri *requestImpl) WithUrl(url string) RequestWrapper {
	ri.url = url
	return ri
}

func (ri *requestImpl) WithHeaders(headers http.Header) RequestWrapper {
	ri.headers = headers
	return ri
}

func (ri *requestImpl) Perform() (*http.Response, error) {
	var resp *http.Response
	var err error

	req, err := ri.builder(ri)
	if req == nil {
		return resp, errors.New("invalid request to perform")
	}
	if err != nil {
		return resp, errors.Wrap(err, "failed to build http request")
	}

	client := http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
