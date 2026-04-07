package app

import (
	productHandler "go-minimal/internal/modules/products/handler"
	productRepo "go-minimal/internal/modules/products/repository"
	productService "go-minimal/internal/modules/products/service"

	"github.com/jackc/pgx/v5"
)

func (a *App) initProducts(db *pgx.Conn) {
	repo := productRepo.NewProductRepo(db)
	svc := productService.NewProductService(repo)
	a.ProductsHandler = productHandler.NewProductHandler(svc)
}
