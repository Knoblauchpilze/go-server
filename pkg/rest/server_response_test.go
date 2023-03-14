package rest

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var dummyID = "eb10f542-c2a8-11ed-befe-18c04d0e6a41"

func TestServerResponse_OK(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	mrw := mockResponseWriter{}

	resp := NewSuccessResponse(id)
	resp.Write(&mrw)

	var expected = fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\"}", dummyID, StatusOK)
	assert.Equal(expected, string(mrw.data))
}

func TestServerResponse_NOK(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	mrw := mockResponseWriter{}

	resp := NewErrorResponse(id)
	resp.Write(&mrw)

	var expected = fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\"}", dummyID, StatusNOK)
	assert.Equal(expected, string(mrw.data))
}

func TestServerResponse_WithDescription(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	mrw := mockResponseWriter{}

	resp := NewSuccessResponse(id)
	resp.WithDescription("haha")
	resp.Write(&mrw)

	expected := fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\",\"Description\":\"haha\"}", dummyID, StatusOK)
	assert.Equal(expected, string(mrw.data))
}

func TestServerResponse_WithDetails(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	mrw := mockResponseWriter{}

	resp := NewSuccessResponse(id)
	resp.WithDetails(23)
	resp.Write(&mrw)

	expected := fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\",\"Details\":23}", dummyID, StatusOK)
	assert.Equal(expected, string(mrw.data))

	test := foo{
		Bar: "haha",
		Baz: -23,
	}
	resp.WithDetails(test)
	resp.Write(&mrw)

	expected = fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\",\"Details\":{\"Bar\":\"haha\",\"Baz\":-23}}", dummyID, StatusOK)
	assert.Equal(expected, string(mrw.data))
}

func TestServerResponse_Pass(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	mrw := mockResponseWriter{}

	resp := NewErrorResponse(id)
	resp.Pass()
	resp.Write(&mrw)

	expected := fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\"}", dummyID, StatusOK)
	assert.Equal(expected, string(mrw.data))
}

func TestServerResponse_Fail(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	mrw := mockResponseWriter{}

	resp := NewSuccessResponse(id)
	resp.Fail()
	resp.Write(&mrw)

	expected := fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\"}", dummyID, StatusNOK)
	assert.Equal(expected, string(mrw.data))
}

func TestServerResponse_WithCode(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	mrw := mockResponseWriter{}

	resp := NewSuccessResponse(id)
	resp.WithCode(http.StatusOK)
	resp.Write(&mrw)

	expected := fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\"}", dummyID, StatusOK)
	assert.Equal(expected, string(mrw.data))
	assert.Equal(http.StatusOK, mrw.code)

	resp.WithCode(http.StatusTeapot)
	resp.Write(&mrw)

	expected = fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\"}", dummyID, StatusNOK)
	assert.Equal(expected, string(mrw.data))
	assert.Equal(http.StatusTeapot, mrw.code)
}

func TestServerResponse_WithCodeAndDescription(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	mrw := mockResponseWriter{}
	code := http.StatusForbidden

	resp := NewSuccessResponse(id)
	resp.WithCodeAndDescription(code)
	resp.Write(&mrw)

	expected := fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\",\"Description\":\"%v\"}", dummyID, StatusNOK, http.StatusText(code))
	assert.Equal(expected, string(mrw.data))
	assert.Equal(http.StatusForbidden, mrw.code)
}

func TestServerResponse_Write(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	mrw := mockResponseWriter{}

	resp := NewSuccessResponse(id)
	resp.Write(&mrw)

	expected := fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\"}", dummyID, StatusOK)
	assert.Equal(expected, string(mrw.data))
	assert.Equal(http.StatusOK, mrw.code)

	resp.WithDescription("haha")
	resp.Write(&mrw)

	expected = fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\",\"Description\":\"haha\"}", dummyID, StatusOK)
	assert.Equal(expected, string(mrw.data))
	assert.Equal(http.StatusOK, mrw.code)

	resp.WithDetails(12.2)
	resp.Write(&mrw)

	expected = fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\",\"Description\":\"haha\",\"Details\":12.2}", dummyID, StatusOK)
	assert.Equal(expected, string(mrw.data))
	assert.Equal(http.StatusOK, mrw.code)
}
