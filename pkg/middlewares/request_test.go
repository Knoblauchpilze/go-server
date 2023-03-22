package middlewares

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestRequestData_FailWithErrorAndCode(t *testing.T) {
	assert := assert.New(t)

	code := http.StatusBadGateway
	inErr := errors.New("haha")
	expectedBody, err := json.Marshal(inErr)
	assert.Nil(err)

	rq := NewRequestData()
	mrw := mockResponseWriter{}
	rq.FailWithErrorAndCode(inErr, code, &mrw)

	resp, err := unmarshalExpectedResponseBody(mrw.data)
	assert.Nil(err)
	assert.Equal(mrw.code, code)
	assert.Equal(resp.Status, "ERROR")
	assert.Equal(string(resp.Details), string(expectedBody))
}

func TestRequestData_WriteDetails(t *testing.T) {
	assert := assert.New(t)

	rq := NewRequestData()
	mrw := mockResponseWriter{}
	rq.WriteDetails(32, &mrw)

	resp, err := unmarshalExpectedResponseBody(mrw.data)
	assert.Nil(err)

	assert.Equal(mrw.code, http.StatusOK)
	assert.Equal(resp.Status, "SUCCESS")

	var val int
	err = json.Unmarshal(resp.Details, &val)
	assert.Nil(err)
	assert.Equal(val, 32)
}
