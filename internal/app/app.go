package app

import (
	userLoginHandler "go-minimal/internal/modules/auth/login/handler"
	"go-minimal/internal/modules/materials/handler"
	userHandler "go-minimal/internal/modules/users/handler"

	"github.com/jackc/pgx/v5"
)

type App struct {
	MaterialHandler  *handler.MaterialHandler
	UserHandler      *userHandler.UserHandler
	UserLoginHandler *userLoginHandler.UserLoginHandler
}

func NewApp(db *pgx.Conn) *App {
	a := &App{}
	a.initMaterials(db)
	a.initUsers(db)
	return a
}
