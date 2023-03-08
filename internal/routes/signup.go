package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
)

var SignUpURLRoute = "/signup"

type signUpRequestDataKeyType string

var signUpRequestDataKey signUpRequestDataKeyType = "signupData"

var userHeaderKey = "User"
var passwordHeaderKey = "Password"

type signUpRequest struct {
	User     string
	Password string
}

func buildSignUpDataFromRequest(w http.ResponseWriter, r *http.Request) (signUpRequest, bool) {
	var user, password string
	var err error

	user, err = rest.GetSingleHeaderFromRequest(r, userHeaderKey)
	if err != nil {
		http.Error(w, "no user provided in sign up request", http.StatusBadRequest)
		return signUpRequest{}, false
	}

	password, err = rest.GetSingleHeaderFromRequest(r, passwordHeaderKey)
	if err != nil {
		http.Error(w, "no password provided in sign up request", http.StatusBadRequest)
		return signUpRequest{}, false
	}

	req := signUpRequest{
		User:     user,
		Password: password,
	}

	return req, true
}

func signUpCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signUpData, ok := buildSignUpDataFromRequest(w, r)
		if !ok {
			return
		}

		ctx := context.WithValue(r.Context(), signUpRequestDataKey, signUpData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SignUpRouter(udb users.UserDb) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(signUpCtx)
		r.Get("/", generateSignUpHandler(udb))
	})

	return r
}

func generateSignUpHandler(udb users.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data, ok := ctx.Value(signUpRequestDataKey).(signUpRequest)
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		id, err := udb.AddUser(data.User, data.Password)
		if err != nil {
			errCode := http.StatusBadRequest
			if err == users.ErrUserCreationFailure {
				errCode = http.StatusInternalServerError
			}

			http.Error(w, fmt.Sprintf("%v", err), errCode)
			return
		}

		rest.SetupStringResponse(w, "{\"user\":\"%s\"}\n", id)
	}
}
