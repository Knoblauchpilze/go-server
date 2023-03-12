package rest

import "fmt"

// body
var ErrFailedToGetBody = fmt.Errorf("failed to get request body")
var ErrBodyParsingFailed = fmt.Errorf("failed to parse request body")

// header
var ErrNoSuchHeader = fmt.Errorf("no such header in request")
var ErrNonUniqueHeader = fmt.Errorf("header is defined multiple times in request")

// response
var ErrInvalidResponse = fmt.Errorf("invalid nil response")
var ErrResponseIsError = fmt.Errorf("response returned error code")
