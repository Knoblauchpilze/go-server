package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ServerResponse interface {
	Pass() ServerResponse
	Fail() ServerResponse

	WithDetails(details interface{}) ServerResponse
	WithCode(httpCode int) ServerResponse

	Write(w http.ResponseWriter)
}

type serverResponseImpl struct {
	RequestID uuid.UUID
	Status    string
	Details   json.RawMessage `json:",omitempty"`
	code      int
}

var StatusOK = "SUCCESS"
var StatusNOK = "ERROR"

func NewSuccessResponse(id uuid.UUID) ServerResponse {
	return &serverResponseImpl{
		RequestID: id,
		Status:    StatusOK,
		code:      http.StatusOK,
	}
}

func NewErrorResponse(id uuid.UUID) ServerResponse {
	return &serverResponseImpl{
		RequestID: id,
		Status:    StatusNOK,
		code:      http.StatusOK,
	}
}

func (sr *serverResponseImpl) Pass() ServerResponse {
	sr.Status = StatusOK
	return sr
}

func (sr *serverResponseImpl) Fail() ServerResponse {
	sr.Status = StatusNOK
	return sr
}

func (sr *serverResponseImpl) WithDetails(details interface{}) ServerResponse {
	var out []byte
	var err error

	// Handle error interface.
	if inErr, ok := details.(error); ok {
		out, err = json.Marshal(inErr.Error())
	} else {
		out, err = json.Marshal(details)
	}

	if err != nil {
		logrus.Errorf("Failed to add details %v to response (%v)", details, err)
	} else {
		sr.Details = out
	}

	return sr
}

func (sr *serverResponseImpl) WithCode(httpCode int) ServerResponse {
	sr.code = httpCode
	if sr.code != http.StatusOK {
		return sr.Fail()
	}

	return sr.Pass()
}

func (sr *serverResponseImpl) Write(w http.ResponseWriter) {
	out, err := json.Marshal(sr)
	if err != nil {
		logrus.Errorf("Failed to setup server response (%v)", err)
	}

	w.WriteHeader(sr.code)
	w.Write(out)
}
