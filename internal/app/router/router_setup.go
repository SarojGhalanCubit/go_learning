package router

import (
	"go-minimal/internal/app"
	authRoute "go-minimal/internal/modules/auth/routes"
	categoriesRoute "go-minimal/internal/modules/categories/routes"
	colorRoute "go-minimal/internal/modules/colors/routes"
	materialsRoute "go-minimal/internal/modules/materials/routes"
	productsRoute "go-minimal/internal/modules/products/routes"
	sizeRoutes "go-minimal/internal/modules/sizes/routes"
	usersRoute "go-minimal/internal/modules/users/routes"
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
		authRoute.RegisterRoutes(r, a.AuthHandler)
	})

	r.Route("/api/v1/materials", func(r chi.Router) {
		materialsRoute.RegisterRoutes(r, a.MaterialHandler)
	})

	r.Route("/api/v1/users", func(r chi.Router) {
		usersRoute.RegisterRoutes(r, a.UserHandler)
	})

	r.Route("/api/v1/colors", func(r chi.Router) {
		colorRoute.RegisterRoutes(r, a.ColorHandler)
	})

	r.Route("/api/v1/sizes", func(r chi.Router) {
		sizeRoutes.RegisterRoutes(r, a.SizeHandler)
	})

	r.Route("/api/v1/categories", func(r chi.Router) {
		categoriesRoute.RegisterRoutes(r, a.CategoriesHandler)
	})

	r.Route("/api/v1/products", func(r chi.Router) {
		productsRoute.RegisterRoutes(r, a.ProductsHandler)
	})
	return &Router{mux: r}
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
