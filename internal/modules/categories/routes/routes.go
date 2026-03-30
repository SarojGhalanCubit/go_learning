package categoriesRoute

import (
	"go-minimal/internal/constants"
	categoriesHandler "go-minimal/internal/modules/categories/handler"

	myMw "go-minimal/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *categoriesHandler.CategoriesHandler) {
	r.Use(myMw.AuthMiddleware)
	r.Group(func(r chi.Router) {

		r.Use(myMw.RequireRole(constants.AdminID, constants.ManagerID))

		r.Get("/getAll", h.GetAllCategories)
		// r.Get("/{id}/getById", h.GeySizeByID)
	})
	// r.Group(func(r chi.Router) {
	// 	r.Use(myMw.RequireRole(constants.AdminID))
	// 	r.Post("/create", h.CreateSize)
	// 	r.Delete("/{id}/delete", h.DeleteSizeByID)
	// 	r.Put("/{id}/update", h.UpdateSize)
	// })
}
