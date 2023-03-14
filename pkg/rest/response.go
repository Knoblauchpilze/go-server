package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GetBodyFromHttpResponseAs(resp *http.Response, out interface{}) error {
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
