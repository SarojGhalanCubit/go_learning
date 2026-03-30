package app

import (
	categoriesHandler "go-minimal/internal/modules/categories/handler"
	categoriesRepo "go-minimal/internal/modules/categories/repository"
	categoriesService "go-minimal/internal/modules/categories/service"

	"github.com/jackc/pgx/v5"
)

func (a *App) initCategories(db *pgx.Conn) {
	repo := categoriesRepo.NewCategoriesRepo(db)
	svc := categoriesService.NewCategoriesService(repo)
	a.CategoriesHandler = categoriesHandler.NewCategoriesHandler(svc)
}
