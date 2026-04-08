package app

import (
	categoriesRepo "go-minimal/internal/modules/categories/repository"
	materialRepository "go-minimal/internal/modules/materials/repository"
	productHandler "go-minimal/internal/modules/products/handler"
	productRepo "go-minimal/internal/modules/products/repository"
	productService "go-minimal/internal/modules/products/service"

	"github.com/jackc/pgx/v5"
)

func (a *App) initProducts(db *pgx.Conn) {
	catRepo := categoriesRepo.NewCategoriesRepo(db)
	matRepo := materialRepository.NewMaterialRepository(db)
	repo := productRepo.NewProductRepo(db)
	svc := productService.NewProductService(repo, catRepo, matRepo)
	a.ProductsHandler = productHandler.NewProductHandler(svc)
}
