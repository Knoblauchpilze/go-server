package rest

import (
	"fmt"
	"net/http"
)

var genericMarshalDataError = "Failed to marshal user data"

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
