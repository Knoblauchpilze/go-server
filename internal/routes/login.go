package routes

import (
	"context"
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

func loginCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var loginData types.UserData

		if err := rest.GetBodyFromHttpRequestAs(r, &loginData); err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), loginRequestDataKey, loginData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LoginRouter(udb users.UserDb, tokens auth.Auth) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(requestCtx, loginCtx)
		r.Post("/", generateLoginHandler(udb, tokens))
	})

	return r
}

func generateLoginHandler(udb users.UserDb, tokens auth.Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := buildServerResponseFromHttpRequest(r)

		ctx := r.Context()
		data, ok := ctx.Value(loginRequestDataKey).(types.UserData)
		if !ok {
			resp.WithCodeAndDescription(http.StatusUnprocessableEntity)
			resp.Write(w)
			return
		}

		user, err := udb.GetUserFromName(data.Name)
		if err != nil {
			resp.WithCodeAndDescription(http.StatusBadRequest)
			resp.WithDetails(err)
			resp.Write(w)
			return
		}

		if user.Password != data.Password {
			resp.WithCodeAndDescription(http.StatusUnauthorized)
			resp.WithDetails("wrong password provided")
			resp.Write(w)
			return
		}

		token, err := tokens.GenerateToken(user.ID, user.Password)
		if err != nil {
			if err == auth.ErrTokenAlreadyExists {
				err = ErrAlreadyLoggedIn
			}

			resp.WithCodeAndDescription(http.StatusBadRequest)
			resp.WithDetails(err)
			resp.Write(w)
			return
		}

		out := types.LoginResponse{
			Token: token,
		}
		resp.WithDetails(out)
		resp.Write(w)
	}
}
