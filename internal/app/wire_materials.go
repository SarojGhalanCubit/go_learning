package app

import (
	"go-minimal/internal/modules/materials/handler"
	"go-minimal/internal/modules/materials/repository"
	"go-minimal/internal/modules/materials/service"

	"github.com/jackc/pgx/v5"
)

func (a *App) initMaterials(db *pgx.Conn) {
	repo := repository.NewMaterialRepository(db)
	svc := service.NewMaterialService(repo)
	a.MaterialHandler = handler.NewMaterialHandler(svc)
}
