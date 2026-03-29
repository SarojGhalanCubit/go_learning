package router

import (
	"go-minimal/internal/app"
	materials "go-minimal/internal/modules/materials/routes"
	users "go-minimal/internal/modules/users/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	mux *chi.Mux
}

func NewRouter(a *app.App) *Router {
	r := chi.NewRouter()
	registerMiddleware(r)

	r.Route("/api/v1/login", func(r chi.Router) {
		materials.RegisterRoutes(r, a.MaterialHandler)
	})

	r.Route("/api/v1/materials", func(r chi.Router) {
		materials.RegisterRoutes(r, a.MaterialHandler)
	})

	r.Route("/api/v1/users", func(r chi.Router) {
		users.RegisterRoutes(r, a.UserHandler)
	})

	return &Router{mux: r}
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
