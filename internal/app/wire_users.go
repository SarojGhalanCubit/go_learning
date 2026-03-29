package app

import (
	"go-minimal/internal/modules/users/handler"
	"go-minimal/internal/modules/users/repository"
	"go-minimal/internal/modules/users/service"

	"github.com/jackc/pgx/v5"
)

func (a *App) initUsers(db *pgx.Conn) {
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	a.UserHandler = handler.NewUserHandler(svc)
}
