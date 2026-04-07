package products

import (
	"go-minimal/internal/constants"
	"go-minimal/internal/middleware"
	productHandler "go-minimal/internal/modules/products/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *productHandler.ProductHandler) {
	r.Use(middleware.AuthMiddleware)

	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireRole(constants.AdminID, constants.ManagerID))

		r.Get("/getAll", h.GetAllProducts)
		r.Get("/create", h.CreateProduct)
	})
}
