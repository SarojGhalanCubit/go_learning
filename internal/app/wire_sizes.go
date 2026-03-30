package app

import (
	sizeHandler "go-minimal/internal/modules/sizes/handler"
	sizeRepository "go-minimal/internal/modules/sizes/repository"
	sizeService "go-minimal/internal/modules/sizes/service"

	"github.com/jackc/pgx/v5"
)

func (a *App) intiSizes(db *pgx.Conn) {
	repo := sizeRepository.NewSizeRepository(db)
	svc := sizeService.NewSizeService(repo)
	a.SizeHandler = sizeHandler.NewSizeHandler(svc)
}
