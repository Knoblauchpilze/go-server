package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
)

// Some inspiration here:
// https://mattermost.com/blog/how-to-build-an-authentication-microservice-in-golang-from-scratch/
// This is why we expect the email and password to be provided as headers.

var LoginURLRoute = "/login"

var loginRequestDataKey stringDataKeyType = "loginData"

var ErrAlreadyLoggedIn = fmt.Errorf("user already logged in")

func buildLoginDataFromRequest(w http.ResponseWriter, r *http.Request) (types.UserData, bool) {
	var data types.UserData
	var err error

	if err = rest.GetBodyFromRequestAs(r, &data); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	}

	return data, err == nil
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

func LoginRouter(udb users.UserDb, tokens auth.Auth) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(loginCtx)
		r.Post("/", generateLoginHandler(udb, tokens))
	})

	return r
}

func generateLoginHandler(udb users.UserDb, tokens auth.Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data, ok := ctx.Value(loginRequestDataKey).(types.UserData)
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		user, err := udb.GetUserFromName(data.Name)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		if user.Password != data.Password {
			http.Error(w, "wrong password provided", http.StatusUnauthorized)
			return
		}

		token, err := tokens.GenerateToken(user.ID, user.Password)
		if err != nil {
			if err == auth.ErrTokenAlreadyExists {
				err = ErrAlreadyLoggedIn
			}

			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		out, err := json.Marshal(token)
		if err != nil {
			rest.SetupInternalErrorResponseWithCause(w, err)
			return
		}

		w.Write(out)
	}
}
