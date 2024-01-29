package main

import (
	"api-ent/db"
	"api-ent/router"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	db.OpenDb()
	router.RegisterRoutes(r)

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
