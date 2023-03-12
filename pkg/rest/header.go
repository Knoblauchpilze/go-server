package rest

import (
	"net/http"
)

func GetHeaderFromRequest(req *http.Request, headerKey string) ([]string, error) {
	header, ok := req.Header[headerKey]
	if !ok {
		return nil, ErrNoSuchHeader
	}

	return header, nil
}

func GetSingleHeaderFromRequest(req *http.Request, headerKey string) (string, error) {
	header, err := GetHeaderFromRequest(req, headerKey)
	if err != nil {
		return "", err
	}

	if len(header) > 1 {
		return "", ErrNonUniqueHeader
	}

	return header[0], nil
}
