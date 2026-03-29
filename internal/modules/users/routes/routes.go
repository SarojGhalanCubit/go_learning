package users

import (
	"github.com/go-chi/chi/v5"
	"go-minimal/internal/constants"
	myMw "go-minimal/internal/middleware"
	"go-minimal/internal/modules/users/handler"
)

func RegisterRoutes(r chi.Router, h *handler.UserHandler) {
	// Public or Generic routes
	r.Get("/getAll", h.GetUsers)
	r.Get("/{id}/getById", h.GetUserByID)

	// Protected Group
	r.Group(func(r chi.Router) {
		r.Use(myMw.AuthMiddleware)
		r.Use(myMw.RequireRole(constants.AdminID))

		r.Post("/create", h.CreateUser)
		r.Put("/{id}/update", h.UpdateUser)
		r.Delete("/{id}/delete", h.DeleteUser)
	})
}
