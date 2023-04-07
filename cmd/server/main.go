package main

import (
	"net/http"

	"github.com/KnoblauchPilze/go-server/internal/routes"
	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	udb := users.NewUserManager()
	tokens := auth.NewAuthenticater()

	r.Mount(routes.SignUpURLRoute, routes.SignUpRouter(udb))
	r.Mount(routes.LoginURLRoute, routes.LoginRouter(udb, tokens))
	r.Mount(routes.UsersURLRoute, routes.UsersRouter(udb, tokens))

	logrus.Infof("Starting server on port 3000...")
	http.ListenAndServe(":3000", r)
}
