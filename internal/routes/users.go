package routes

import (
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/KnoblauchPilze/go-server/pkg/middlewares"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var UsersURLRoute = "/users"

var userIdDataKey = "user"

func UsersRouter(udb users.UserDb, tokens auth.Auth) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(middlewares.RequestCtx, generateAuthenticationContext(tokens))
		r.Get("/", generateListUsersHandler(udb))

		r.Route("/{user}", func(r chi.Router) {
			r.Get("/", generateUsersHandler(udb))
		})
	})

	return r
}

func generateListUsersHandler(udb users.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, ok := middlewares.GetRequestDataFromContextOrFail(w, r)
		if !ok {
			return
		}

		people := udb.GetUsers()
		reqData.WriteDetails(people, w)
	}
}

func generateUsersHandler(udb users.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, ok := middlewares.GetRequestDataFromContextOrFail(w, r)
		if !ok {
			return
		}

		id, err := getUserIdFromRequest(r)
		if err != nil {
			reqData.FailWithErrorAndCode(err, http.StatusBadRequest, w)
			return
		}

		ud, err := udb.GetUser(id)
		if err != nil {
			reqData.FailWithErrorAndCode(err, http.StatusBadRequest, w)
			return
		}

		reqData.WriteDetails(ud, w)
	}
}

func getUserIdFromRequest(r *http.Request) (uuid.UUID, error) {
	var err error
	var id uuid.UUID

	qp := chi.URLParam(r, userIdDataKey)
	if len(qp) == 0 {
		return id, errors.New("no user Id provided")
	}

	id, err = uuid.Parse(qp)
	if err != nil {
		return id, errors.Wrap(err, "invalid user id provided")
	}

	return id, nil
}
