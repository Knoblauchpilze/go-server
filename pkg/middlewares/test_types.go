package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/google/uuid"
)

var errSomeError = fmt.Errorf("some error")

type mockResponseWriter struct {
	code int
	data []byte
}

func (mrw *mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (mrw *mockResponseWriter) Write(out []byte) (int, error) {
	mrw.data = out
	return len(mrw.data), nil
}

func (mrw *mockResponseWriter) WriteHeader(statusCode int) {
	mrw.code = statusCode
}

type mockAuth struct {
	isError    bool
	token      string
	expiration time.Time
}

func (ma mockAuth) GenerateToken(user uuid.UUID, password string) (auth.Token, error) {
	if ma.isError {
		return auth.Token{}, errSomeError
	}

	t := auth.Token{
		User:       user,
		Expiration: ma.expiration,
		Value:      ma.token,
	}

	return t, nil
}

func (ma mockAuth) GetToken(user uuid.UUID) (auth.Token, error) {
	return ma.GenerateToken(user, "hihi")
}

type expectedResponseBody struct {
	RequestId uuid.UUID
	Status    string
	Details   json.RawMessage
}

func unmarshalExpectedResponseBody(body []byte) (expectedResponseBody, error) {
	var out expectedResponseBody
	err := json.Unmarshal(body, &out)
	return out, err
}

func defaultHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rd, res := GetRequestDataFromContextOrFail(w, r)
		if res {
			rd.WriteDetails(res, w)
		}
	})
}
