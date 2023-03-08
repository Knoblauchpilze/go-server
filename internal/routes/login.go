package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Some inspiration here:
// https://mattermost.com/blog/how-to-build-an-authentication-microservice-in-golang-from-scratch/
// This is why we expect the email and password to be provided as headers.

var LoginURLRoute = "/login"

type loginRequestDataKeyType string

var loginRequestDataKey loginRequestDataKeyType = "loginData"

type loginRequest struct {
	User     string
	Password string
}

func buildLoginDataFromRequest(w http.ResponseWriter, r *http.Request) (loginRequest, bool) {
	var user, password string
	var err error

	user, err = rest.GetSingleHeaderFromRequest(r, userHeaderKey)
	if err != nil {
		http.Error(w, "no user provided in login request", http.StatusBadRequest)
		return loginRequest{}, false
	}

	password, err = rest.GetSingleHeaderFromRequest(r, passwordHeaderKey)
	if err != nil {
		http.Error(w, "no password provided in login request", http.StatusBadRequest)
		return loginRequest{}, false
	}

	req := loginRequest{
		User:     user,
		Password: password,
	}

	return req, true
}

func loginCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginData, ok := buildLoginDataFromRequest(w, r)
		if !ok {
			return
		}

		ctx := context.WithValue(r.Context(), loginRequestDataKey, loginData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LoginRouter(udb users.UserDb) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(loginCtx)
		r.Post("/", generateLoginHandler(udb))
	})

	return r
}

func generateLoginHandler(udb users.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data, ok := ctx.Value(loginRequestDataKey).(loginRequest)
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		user, err := udb.GetUserFromName(data.User)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		if user.Password != data.Password {
			http.Error(w, "wrong password provided", http.StatusUnauthorized)
			return
		}

		token, err := uuid.NewUUID()
		if err != nil {
			rest.SetupInternalErrorResponseWithCause(w, err)
			return
		}

		rest.SetupStringResponse(w, "{\"token\":\"%s\"}\n", token)
	}
}
