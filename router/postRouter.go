package router

import (
	"api-ent/controller"
	"api-ent/middleware"

	"github.com/go-chi/chi/v5"
)

func registerPostRouter(r *chi.Mux) {
	r.Route("/post", func(r chi.Router) {
		r.Get("/", controller.GetAllPosts)
		r.Get("/{id}", controller.GetPost)
		r.Get("/title", controller.GetAllTitles)
		r.Get("/title/{title}", controller.GetPostsByTitle)
		r.Route("/", func(r chi.Router) {
			r.Use(middleware.JWT)
			r.Post("/", controller.AddPost)
		})
	})
}
