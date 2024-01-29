package router

import (
	"api-ent/controller"

	"github.com/go-chi/chi/v5"
)

func registerUserRouter(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/", controller.GetAllUsers)
		r.Post("/", controller.AddUser)
		r.Get("/{id}", controller.GetUser)
	})
}
