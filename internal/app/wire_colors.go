package app

import (
	"go-minimal/internal/modules/colors/handler"
	"go-minimal/internal/modules/colors/repository"
	"go-minimal/internal/modules/colors/service"

	"github.com/jackc/pgx/v5"
)

func (a *App) initColors(db *pgx.Conn) {
	repo := colorsRepository.NewColorsRepository(db)
	svc := colorService.NewColorsService(repo)
	a.ColorHandler = colorHandler.NewColorHandler(svc)
}
