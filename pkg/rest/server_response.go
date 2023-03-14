package rest

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ServerResponse interface {
	WithDescription(desc string) ServerResponse
	WithDetails(details interface{}) (ServerResponse, error)

	Succeeds() ServerResponse
	Fails() ServerResponse
}

type serverResponseImpl struct {
	RequestID   uuid.UUID
	Status      string
	Details     json.RawMessage `json:",omitempty"`
	Description string          `json:",omitempty"`
}

var StatusOK = "SUCCESS"
var StatusNOK = "ERROR"

func NewSuccessResponse(id uuid.UUID) ServerResponse {
	return &serverResponseImpl{
		RequestID: id,
		Status:    StatusOK,
	}
}

func NewErrorResponse(id uuid.UUID) ServerResponse {
	return &serverResponseImpl{
		RequestID: id,
		Status:    StatusNOK,
	}
}

func (sr *serverResponseImpl) WithDescription(desc string) ServerResponse {
	sr.Description = desc
	return sr
}

func (sr *serverResponseImpl) WithDetails(details interface{}) (ServerResponse, error) {
	out, err := json.Marshal(details)
	if err != nil {
		logrus.Errorf("Failed to add details %v to response (%v)", details, err)
	} else {
		sr.Details = out
	}

	return sr, err
}

func (sr *serverResponseImpl) Succeeds() ServerResponse {
	sr.Status = StatusOK
	return sr
}

func (sr *serverResponseImpl) Fails() ServerResponse {
	sr.Status = StatusNOK
	return sr
}
