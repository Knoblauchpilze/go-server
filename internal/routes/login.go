package routes

import (
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
)

// Some inspiration here:
// https://mattermost.com/blog/how-to-build-an-authentication-microservice-in-golang-from-scratch/
// This is why we expect the email and password to be provided as headers.

var LoginURLRoute = "/login"

func LoginRouter(udb users.UserDb, tokens auth.Auth) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(requestCtx)
		r.Post("/", generateLoginHandler(udb, tokens))
	})

	return r
}

func generateLoginHandler(udb users.UserDb, tokens auth.Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, ok := getRequestDataFromContextOrFail(w, r)
		if !ok {
			return
		}

		var err error
		var in types.UserData

		if in, err = getUserDataFromRequest(r); err != nil {
			reqData.failWithErrorAndCode(err, http.StatusBadRequest, w)
			return
		}

		user, err := udb.GetUserFromName(in.Name)
		if err != nil {
			reqData.failWithErrorAndCode(err, http.StatusBadRequest, w)
			return
		}

		out, err := loginUser(in, user, udb, tokens)
		if err != nil {
			err = interpretLoginFailure(err)
			reqData.failWithErrorAndCode(err, http.StatusBadRequest, w)
			return
		}

		reqData.writeDetails(out, w)
	}
}

func loginUser(in types.UserData, ud users.User, udb users.UserDb, tokens auth.Auth) (types.LoginResponse, error) {
	var err error
	var out types.LoginResponse

	if in.Password != ud.Password {
		return out, errors.New("wrong password provided")
	}

	out.Token, err = tokens.GenerateToken(ud.ID, ud.Password)
	return out, err
}

func interpretLoginFailure(err error) error {
	if errors.IsErrorWithCode(err, errors.ErrTokenAlreadyExists) {
		return errors.Wrap(err, "user already logged in")
	}

	return err
}
