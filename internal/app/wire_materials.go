package app

import (
	"go-minimal/internal/modules/materials/handler"
	"go-minimal/internal/modules/materials/repository"
	"go-minimal/internal/modules/materials/service"

	"github.com/jackc/pgx/v5"
)

func (a *App) initMaterials(db *pgx.Conn) {
	repo := materialRepository.NewMaterialRepository(db)
	svc := materialService.NewMaterialService(repo)
	a.MaterialHandler = materialsHandler.NewMaterialHandler(svc)
}
