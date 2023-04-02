package connection

import "net/http"

type RequestWrapper interface {
	Perform() (*http.Response, error)
}
