package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var ErrFailedToGetBody = fmt.Errorf("failed to get request body")
var ErrBodyParsingFailed = fmt.Errorf("failed to parse request body")

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
