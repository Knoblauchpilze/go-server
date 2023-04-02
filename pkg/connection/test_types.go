package connection

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var errSomeError = fmt.Errorf("some error")

type mockHttpClient struct {
	inReq         *http.Request
	expectedResp  *http.Response
	expectedError error
}

func (mc *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	mc.inReq = req
	return mc.expectedResp, mc.expectedError
}

func generateHttpResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusAlreadyReported,
		Header: http.Header{
			"haha": []string{"gigi"},
		},
		Body: io.NopCloser(bytes.NewReader([]byte("some data"))),
	}
}
