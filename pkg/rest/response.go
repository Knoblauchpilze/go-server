package rest

import (
	"fmt"
	"net/http"
)

func SetupStringResponse(w http.ResponseWriter, format string, a ...any) {
	out := []byte(fmt.Sprintf(format, a...))
	w.Write(out)
}
