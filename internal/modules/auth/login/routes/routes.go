package loginRoute

import (
	userLoginHandler "go-minimal/internal/modules/auth/login/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *userLoginHandler.UserLoginHandler) {
	r.Post("/", h.Login)
}
