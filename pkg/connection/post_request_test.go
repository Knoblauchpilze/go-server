package connection

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpPostRequestBuilder_Fail(t *testing.T) {
	assert := assert.New(t)

	mc := &mockHttpClient{
		expectedError: errSomeError,
		expectedResp:  nil,
	}
	headers := http.Header{
		"hihi": []string{"haha"},
	}

	rb := NewHttpPostRequestBuilder()
	rb.SetUrl("haha")
	rb.SetHeaders(headers)
	rb.setHttpClient(mc)
	rw, err := rb.Build()
	assert.Nil(err)

	resp, err := rw.Perform()
	assert.Equal(mc.expectedResp, resp)
	assert.Equal(mc.expectedError, err)
	assert.Equal(headers, mc.inReq.Header)
}

func TestHttpPostRequestBuilder_Success(t *testing.T) {
	assert := assert.New(t)

	mc := &mockHttpClient{
		expectedError: nil,
		expectedResp: &http.Response{
			StatusCode: http.StatusResetContent,
			Header: http.Header{
				"juju": []string{"koko"},
			},
			Body: io.NopCloser(bytes.NewReader([]byte{32})),
		},
	}

	rb := NewHttpGetRequestBuilder()
	rb.SetUrl("haha")
	rb.SetBody("jaja", 26)
	rb.setHttpClient(mc)
	rw, err := rb.Build()
	assert.Nil(err)

	resp, err := rw.Perform()
	assert.Equal(mc.expectedResp, resp)
	assert.Equal(mc.expectedError, err)
	headers := http.Header{
		"Content-Type": []string{"jaja"},
	}
	assert.Equal(headers, mc.inReq.Header)
	out, err := io.ReadAll(resp.Body)
	assert.Equal([]byte{32}, out)
	assert.Nil(err)
}
