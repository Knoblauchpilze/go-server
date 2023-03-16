package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Response interface {
	Pass() Response
	Fail() Response

	WithDetails(details interface{}) Response
	WithCode(httpCode int) Response

	Write(w http.ResponseWriter)
}

type responseImpl struct {
	RequestID uuid.UUID
	Status    string
	Details   json.RawMessage `json:",omitempty"`
	code      int
}

var StatusOK = "SUCCESS"
var StatusNOK = "ERROR"

func NewSuccessResponse(id uuid.UUID) Response {
	return &responseImpl{
		RequestID: id,
		Status:    StatusOK,
		code:      http.StatusOK,
	}
}

func NewErrorResponse(id uuid.UUID) Response {
	return &responseImpl{
		RequestID: id,
		Status:    StatusNOK,
		code:      http.StatusOK,
	}
}

func (ri *responseImpl) Pass() Response {
	ri.Status = StatusOK
	return ri
}

func (ri *responseImpl) Fail() Response {
	ri.Status = StatusNOK
	return ri
}

func (ri *responseImpl) WithDetails(details interface{}) Response {
	var out []byte
	var err error

	// Handle error interface.
	out, err = json.Marshal(details)

	if err != nil {
		logrus.Errorf("Failed to add details %v to response (%v)", details, err)
	} else {
		ri.Details = out
	}

	return ri
}

func (ri *responseImpl) WithCode(httpCode int) Response {
	ri.code = httpCode
	if ri.code != http.StatusOK {
		return ri.Fail()
	}

	return ri.Pass()
}

func (ri *responseImpl) Write(w http.ResponseWriter) {
	out, err := json.Marshal(ri)
	if err != nil {
		logrus.Errorf("Failed to setup server response (%v)", err)
	}

	w.WriteHeader(ri.code)
	w.Write(out)
}

func GetBodyFromHttpResponseAs(resp *http.Response, out interface{}) error {
	if resp == nil {
		return errors.NewCode(errors.ErrNoResponse)
	}
	if resp.Body == nil {
		if resp.StatusCode != http.StatusOK {
			logrus.Errorf("Response returned code %d (%v)", resp.StatusCode, http.StatusText(resp.StatusCode))
			return errors.NewCode(errors.ErrResponseIsError)
		}
		return errors.NewCode(errors.ErrFailedToGetBody)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.WrapCode(err, errors.ErrFailedToGetBody)
	}

	var in responseImpl
	err = json.Unmarshal(data, &in)
	if err != nil {
		return errors.WrapCode(err, errors.ErrBodyParsingFailed)
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Response returned code %d (%v): %v", resp.StatusCode, http.StatusText(resp.StatusCode), string(in.Details))
		return errors.NewCode(errors.ErrResponseIsError)
	}

	err = json.Unmarshal(in.Details, out)
	if err != nil {
		logrus.Errorf("Failed to parse %v (err: %v)", string(data), err)
		return errors.WrapCode(err, errors.ErrBodyParsingFailed)
	}

	return nil
}
