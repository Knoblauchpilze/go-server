package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GetBodyFromHttpResponseAs(resp *http.Response, out interface{}) error {
	if resp == nil {
		return ErrNoResponse
	}
	if resp.Body == nil {
		if resp.StatusCode != http.StatusOK {
			return ErrResponseIsError
		}
		return ErrFailedToGetBody
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ErrFailedToGetBody
	}

	var in responseImpl
	err = json.Unmarshal(data, &in)
	if err != nil {
		return ErrInvalidResponse
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Response returned code %d (%v): %v", resp.StatusCode, http.StatusText(resp.StatusCode), string(in.Details))
		return ErrResponseIsError
	}

	err = json.Unmarshal(in.Details, out)
	if err != nil {
		logrus.Errorf("Failed to parse %v (err: %v)", string(data), err)
		return ErrBodyParsingFailed
	}

	return nil
}
