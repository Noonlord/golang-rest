package router

import (
	"api-ent/controller"

	"github.com/go-chi/chi/v5"
)

func registerAuthRouter(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", controller.Login)
	})
}
