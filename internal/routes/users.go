package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var UsersURLRoute = "/users"

type userIDDayaKeyType string

var userIDDataKey userIDDayaKeyType = "user"

func userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := chi.URLParam(r, string(userIDDataKey))
		if len(user) == 0 {
			http.Error(w, "no user ID provided", http.StatusBadRequest)
			return
		}

		userID, err := uuid.Parse(user)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), userIDDataKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UsersRouter(udb users.UserDb) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", generateListUsersHandler(udb))

		r.Route("/{user}", func(r chi.Router) {
			r.Use(userCtx)
			r.Get("/", generateUsersHandler(udb))
		})
	})

	return r
}

func generateListUsersHandler(udb users.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := buildServerResponseFromHttpRequest(r)

		people := udb.GetUsers()
		resp.WithDetails(people)
		resp.Write(w)
	}
}

func generateUsersHandler(udb users.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := buildServerResponseFromHttpRequest(r)

		ctx := r.Context()
		id, ok := ctx.Value(userIDDataKey).(uuid.UUID)
		if !ok {
			resp.WithCodeAndDescription(http.StatusUnprocessableEntity)
			resp.Write(w)
			return
		}

		userData, err := udb.GetUser(id)
		if err != nil {
			resp.WithCodeAndDescription(http.StatusBadRequest)
			resp.WithDetails(err)
			resp.Write(w)
			return
		}

		resp.WithDetails(userData)
		resp.Write(w)
	}
}
