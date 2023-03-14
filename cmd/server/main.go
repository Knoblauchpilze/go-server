package main

// TODO: Common response from the server with an error.

import (
	"fmt"
	"net/http"

	"github.com/KnoblauchPilze/go-server/internal/routes"
	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	udb := users.NewUserDb()
	tokens := auth.NewAuth()

	r.Mount(routes.SignUpURLRoute, routes.SignUpRouter(udb))
	r.Mount(routes.LoginURLRoute, routes.LoginRouter(udb, tokens))
	r.Mount(routes.UsersURLRoute, routes.UsersRouter(udb))

	fmt.Printf("Starting server on port 3000...\n")
	http.ListenAndServe(":3000", r)
}
