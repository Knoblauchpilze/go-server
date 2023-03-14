package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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

func generateResponseWithBody(body []byte) *http.Response {
	resp := http.Response{
		StatusCode: http.StatusOK,
	}

	rdr := bytes.NewReader(body)
	resp.Body = io.NopCloser(rdr)

	return &resp
}

type mockResponseWriter struct {
	code int
	data []byte
}

func (mrw *mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (mrw *mockResponseWriter) Write(out []byte) (int, error) {
	mrw.data = out
	return len(mrw.data), nil
}

func (mrw *mockResponseWriter) WriteHeader(statusCode int) {
	mrw.code = statusCode
}
