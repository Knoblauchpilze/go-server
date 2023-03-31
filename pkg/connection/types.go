package connection

import "net/http"

type Request interface {
	WithUrl(url string) Request
	WithHeaders(headers http.Header) Request
	Perform() (*http.Response, error)
}
