package app

import (
	authHandler "go-minimal/internal/modules/auth/handler"
	colorHandler "go-minimal/internal/modules/colors/handler"
	"go-minimal/internal/modules/materials/handler"
	userHandler "go-minimal/internal/modules/users/handler"

	"github.com/jackc/pgx/v5"
)

type App struct {
	MaterialHandler *materialsHandler.MaterialHandler
	UserHandler     *userHandler.UserHandler
	AuthHandler     *authHandler.AuthHandler
	ColorHandler    *colorHandler.ColorHandler
}

func NewApp(db *pgx.Conn) *App {
	a := &App{}
	a.initMaterials(db)
	a.initUsers(db)
	a.initAuth(db)
	a.initColors(db)
	return a
}
