package app

import (
	authHandler "go-minimal/internal/modules/auth/handler"
	categoriesHandler "go-minimal/internal/modules/categories/handler"
	colorHandler "go-minimal/internal/modules/colors/handler"
	"go-minimal/internal/modules/materials/handler"
	sizeHandler "go-minimal/internal/modules/sizes/handler"
	userHandler "go-minimal/internal/modules/users/handler"

	"github.com/jackc/pgx/v5"
)

type App struct {
	MaterialHandler   *materialsHandler.MaterialHandler
	UserHandler       *userHandler.UserHandler
	AuthHandler       *authHandler.AuthHandler
	ColorHandler      *colorHandler.ColorHandler
	SizeHandler       *sizeHandler.SizeHandler
	CategoriesHandler *categoriesHandler.CategoriesHandler
}

func NewApp(db *pgx.Conn) *App {
	a := &App{}
	a.initMaterials(db)
	a.initUsers(db)
	a.initAuth(db)
	a.initColors(db)
	a.intiSizes(db)
	a.initCategories(db)
	return a
}
