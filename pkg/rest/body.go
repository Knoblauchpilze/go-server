package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetBodyFromRequestAs(req *http.Request, out interface{}) error {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return ErrFailedToGetBody
	}

	err = json.Unmarshal(data, out)
	if err != nil {
		return ErrBodyParsingFailed
	}

	return nil
}
