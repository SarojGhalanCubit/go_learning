package app

import (
	userHandler "go-minimal/internal/modules/users/handler"
	"go-minimal/internal/modules/users/repository"
	"go-minimal/internal/modules/users/service"

	"github.com/jackc/pgx/v5"
)

func (a *App) initUsers(db *pgx.Conn) {
	repo := userRepository.NewUserRepository(db)
	svc := userService.NewUserService(repo)
	a.UserHandler = userHandler.NewUserHandler(svc)
}
