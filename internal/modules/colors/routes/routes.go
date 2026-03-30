package colorRoute

import (
	"go-minimal/internal/constants"
	colorHandler "go-minimal/internal/modules/colors/handler"

	myMw "go-minimal/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *colorHandler.ColorHandler) {

	r.Use(myMw.AuthMiddleware)
	r.Group(func(r chi.Router) {
		r.Use(myMw.RequireRole(constants.AdminID, constants.ManagerID))

		r.Get("/getAll", h.GetAllColors)
		r.Get("/{id}/getById", h.GeyByColorID)
	})

	r.Group(func(r chi.Router) {
		r.Use(myMw.RequireRole(constants.AdminID))

		r.Post("/create", h.CreateColor)
		r.Put("/{id}/update", h.UpdateColor)
		r.Delete("/{id}/delete", h.DeleteColorsByID)
	})

}
