package authRoute

import (
	"github.com/go-chi/chi/v5"
	authHandler "go-minimal/internal/modules/auth/handler"
)

func RegisterRoutes(r chi.Router, h *authHandler.AuthHandler) {
	r.Post("/", h.Login)
}
