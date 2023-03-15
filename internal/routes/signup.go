package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
)

var SignUpURLRoute = "/signup"

var signUpRequestDataKey stringDataKeyType = "signupData"

func signUpCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var signUpData types.UserData

		if err := rest.GetBodyFromHttpRequestAs(r, &signUpData); err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), signUpRequestDataKey, signUpData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SignUpRouter(udb users.UserDb) http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(requestCtx, signUpCtx)
		r.Post("/", generateSignUpHandler(udb))
	})

	return r
}

func generateSignUpHandler(udb users.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := buildServerResponseFromHttpRequest(r)

		ctx := r.Context()
		data, ok := ctx.Value(signUpRequestDataKey).(types.UserData)
		if !ok {
			resp.WithCode(http.StatusUnprocessableEntity)
			resp.Write(w)
			return
		}

		id, err := udb.AddUser(data.Name, data.Password)
		if err != nil {
			errCode := http.StatusBadRequest
			if err == users.ErrUserCreationFailure {
				errCode = http.StatusInternalServerError
			}

			resp.WithCode(errCode)
			resp.WithDetails(err)
			resp.Write(w)
			return
		}

		out := types.SignUpResponse{
			ID: id,
		}
		resp.WithDetails(out)
		resp.Write(w)
	}
}
