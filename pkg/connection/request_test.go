package connection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest_Perform_Fail(t *testing.T) {
	assert := assert.New(t)

	rb := newRequestBuilder()
	rb.SetUrl("haha")

	rb.setHttpRequestBuilder(errorHttpRequestBuilder)
	rw, err := rb.Build()
	assert.Nil(err)
	_, err = rw.Perform()
	assert.NotNil(err)

	rb.setHttpRequestBuilder(nilRequestHttpRequestBuilder)
	rw, err = rb.Build()
	assert.Nil(err)
	resp, _ := rw.Perform()
	assert.Nil(resp)
}
