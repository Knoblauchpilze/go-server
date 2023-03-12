package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

var genericMarshalDataError = "Failed to marshal user data"

func GetBodyFromResponseAs(resp *http.Response, out interface{}) error {
	if resp == nil {
		return ErrInvalidResponse
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Response returned code %d (%v)", resp.StatusCode, http.StatusText(resp.StatusCode))
		return ErrResponseIsError
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ErrFailedToGetBody
	}

	err = json.Unmarshal(data, out)
	if err != nil {
		logrus.Errorf("Failed to parse %v (err: %v)", string(data), err)
		return ErrBodyParsingFailed
	}

	return nil
}

func SetupStringResponse(w http.ResponseWriter, format string, a ...any) {
	out := []byte(fmt.Sprintf(format, a...))
	w.Write(out)
}

func SetupInternalErrorResponse(w http.ResponseWriter) {
	out := []byte(genericMarshalDataError)
	w.Write(out)
}

func SetupInternalErrorResponseWithCause(w http.ResponseWriter, cause interface{}) {
	out := []byte(fmt.Sprintf("%s (cause: %v)", genericMarshalDataError, cause))
	w.Write(out)
}
