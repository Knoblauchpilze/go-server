package connection

import (
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
)

func NewGetRequest(url string, headers http.Header) Request {
	return newRequest(url, headers, buildGetRequest)
}

func buildGetRequest(ri *requestImpl) (*http.Request, error) {
	req, err := http.NewRequest("GET", ri.url, nil)
	if err != nil {
		return req, errors.WrapCode(err, errors.ErrGetRequestFailed)
	}

	req.Header = ri.headers

	return req, nil
}
