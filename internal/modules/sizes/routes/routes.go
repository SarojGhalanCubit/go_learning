package sizeRoutes

import (
	"go-minimal/internal/constants"
	sizeHandler "go-minimal/internal/modules/sizes/handler"

	myMw "go-minimal/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *sizeHandler.SizeHandler) {
	r.Use(myMw.AuthMiddleware)
	r.Group(func(r chi.Router) {

		r.Use(myMw.RequireRole(constants.AdminID, constants.ManagerID))

		r.Get("/getAll", h.GetAllSizes)
		r.Get("/{id}/getById", h.GeySizeByID)
	})
	r.Group(func(r chi.Router) {
		r.Use(myMw.RequireRole(constants.AdminID))
		r.Post("/create", h.CreateSize)
		r.Delete("/{id}/delete", h.DeleteSizeByID)
		r.Put("/{id}/update", h.UpdateSize)
	})
}
