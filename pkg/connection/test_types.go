package connection

import (
	"fmt"
	"net/http"
)

var errSomeError = fmt.Errorf("some error")

func nilRequestHttpRequestBuilder(ri *requestImpl) (*http.Request, error) {
	return nil, nil
}

func errorHttpRequestBuilder(ri *requestImpl) (*http.Request, error) {
	return nil, errSomeError
}

type mockHttpClient struct {
	inReq         *http.Request
	expectedResp  *http.Response
	expectedError error
}

func (mc *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	mc.inReq = req
	return mc.expectedResp, mc.expectedError
}
