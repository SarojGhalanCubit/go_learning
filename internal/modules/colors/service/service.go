package colorService

import (
	"context"
	colorModel "go-minimal/internal/modules/colors/model"
	colorsRepository "go-minimal/internal/modules/colors/repository"
)

type ColorsService struct {
	repo colorsRepository.ColorsRepositoryI
}

func NewColorsService(repo colorsRepository.ColorsRepositoryI) *ColorsService {
	return &ColorsService{
		repo: repo,
	}
}

func (s *ColorsService) GetAllColors(ctx context.Context) ([]colorModel.Colors, error) {
	return s.repo.GetAllColors(ctx)
}

func (s *ColorsService) CreateColor(ctx context.Context, color colorModel.CreateColor) (colorModel.Colors, error) {
	return s.repo.CreateColor(ctx, color)
}

func (s *ColorsService) GeyByColorID(ctx context.Context, colorID string) (colorModel.Colors, error) {

	materialFinded, err := s.repo.FIndByColorID(ctx, colorID)

	if err != nil {
		return colorModel.Colors{}, err
	}

	return s.repo.FIndByColorID(ctx, materialFinded.ID.String())
}

func (s *ColorsService) UpdateColor(ctx context.Context, colorID string, color colorModel.CreateColor) (colorModel.Colors, error) {
	colorFinded, err := s.repo.FIndByColorID(ctx, colorID)
	if err != nil {
		return colorModel.Colors{}, err
	}
	return s.repo.UpdateColor(ctx, colorFinded.ID.String(), color)
}

func (s *ColorsService) DeleteColorsByID(ctx context.Context, colorID string) (colorModel.Colors, error) {
	colorFinded, err := s.repo.FIndByColorID(ctx, colorID)

	if err != nil {
		return colorModel.Colors{}, err
	}
	return s.repo.DeleteByColorID(ctx, colorFinded.ID.String())
}
