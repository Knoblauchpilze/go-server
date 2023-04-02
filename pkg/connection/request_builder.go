package connection

import (
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
)

type requestBuilder struct {
	url     string
	headers http.Header
	builder httpRequestBuilder
	client  HttpClient
}

func newRequestBuilder() *requestBuilder {
	return &requestBuilder{
		client: &http.Client{},
	}
}

func (rb *requestBuilder) setUrl(url string) {
	rb.url = url
}

func (rb *requestBuilder) setHeaders(headers http.Header) {
	rb.headers = headers
}

func (rb *requestBuilder) setHttpRequestBuilder(builder httpRequestBuilder) {
	rb.builder = builder
}

func (rb *requestBuilder) setHttpClient(client HttpClient) {
	rb.client = client
}

func (rb *requestBuilder) build() (RequestWrapper, error) {
	if len(rb.url) == 0 {
		return nil, errors.New("invalid url for request")
	}
	if rb.builder == nil {
		return nil, errors.New("invalid builder for request")
	}

	ri := &requestImpl{
		url:     rb.url,
		headers: rb.headers,
		builder: rb.builder,
		client:  rb.client,
	}

	return ri, nil
}
