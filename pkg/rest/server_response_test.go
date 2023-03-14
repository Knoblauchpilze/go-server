package rest

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var dummyID = "eb10f542-c2a8-11ed-befe-18c04d0e6a41"

func TestServerResponse_OK(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	resp := NewSuccessResponse(id)

	out, err := json.Marshal(resp)

	var expected = fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\"}", dummyID, StatusOK)
	assert.Equal(string(out), expected)
	assert.Nil(err)
}

func TestServerResponse_NOK(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	resp := NewErrorResponse(id)

	out, err := json.Marshal(resp)

	var expected = fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\"}", dummyID, StatusNOK)
	assert.Equal(string(out), expected)
	assert.Nil(err)
}

func TestServerResponse_WithDescription(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	resp := NewSuccessResponse(id)
	resp.WithDescription("haha")

	out, err := json.Marshal(resp)

	expected := fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\",\"Description\":\"haha\"}", dummyID, StatusOK)
	assert.Equal(string(out), expected)
	assert.Nil(err)
}

func TestServerResponse_WithDetails(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse(dummyID)
	resp := NewSuccessResponse(id)
	resp.WithDetails(23)
	out, err := json.Marshal(resp)

	expected := fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\",\"Details\":23}", dummyID, StatusOK)
	assert.Equal(string(out), expected)
	assert.Nil(err)

	test := foo{
		Bar: "haha",
		Baz: -23,
	}
	resp.WithDetails(test)
	out, err = json.Marshal(resp)

	expected = fmt.Sprintf("{\"RequestID\":\"%s\",\"Status\":\"%v\",\"Details\":{\"Bar\":\"haha\",\"Baz\":-23}}", dummyID, StatusOK)
	assert.Equal(string(out), expected)
	assert.Nil(err)
}
