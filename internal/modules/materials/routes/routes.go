package materials

import (
	"github.com/go-chi/chi/v5"
	"go-minimal/internal/constants"
	myMw "go-minimal/internal/middleware"
	"go-minimal/internal/modules/materials/handler"
)

func RegisterRoutes(r chi.Router, h *materialsHandler.MaterialHandler) {
	r.Use(myMw.AuthMiddleware)

	r.Group(func(r chi.Router) {
		r.Use(myMw.RequireRole(constants.AdminID, constants.ManagerID))

		r.Get("/getAll", h.GetAllMaterial)
		r.Get("/{id}/getById", h.GeyByMaterialID)
	})

	// Protected Group
	r.Group(func(r chi.Router) {
		r.Use(myMw.RequireRole(constants.AdminID))
		r.Post("/create", h.CreateMaterial)
		r.Put("/{id}/update", h.UpdateMaterial)
		r.Delete("/{id}/delete", h.DeleteMaterial)
	})
}
