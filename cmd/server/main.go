package main

import (
	"fmt"
	"net/http"

	"github.com/KnoblauchPilze/go-server/internal"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there\n"))
	})

	udb := users.NewUserDb()

	r.Mount(internal.LoginURLRoute, internal.LoginRouter(udb))

	fmt.Printf("Starting server on port 3000...\n")
	http.ListenAndServe(":3000", r)
}
