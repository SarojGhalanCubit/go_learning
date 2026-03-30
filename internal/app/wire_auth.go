package app

import (
	authHandler "go-minimal/internal/modules/auth/handler"
	authService "go-minimal/internal/modules/auth/service"
	userRepo "go-minimal/internal/modules/users/repository"

	"github.com/jackc/pgx/v5"
)

func (a *App) initAuth(db *pgx.Conn) {
	repo := userRepo.NewUserRepository(db)
	svc := authService.NewAuthService(repo)
	a.AuthHandler = authHandler.NewAuthHandler(svc)
}
