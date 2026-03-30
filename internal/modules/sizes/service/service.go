package sizeService

import (
	"context"
	sizeModel "go-minimal/internal/modules/sizes/model"
	sizeRepository "go-minimal/internal/modules/sizes/repository"
)

type SizeService struct {
	repo sizeRepository.SizeRepositoryI
}

func NewSizeService(r sizeRepository.SizeRepositoryI) *SizeService {
	return &SizeService{repo: r}
}

func (s *SizeService) GetAllSizes(ctx context.Context) ([]sizeModel.Sizes, error) {
	return s.repo.GetAllSizes(ctx)
}

func (s *SizeService) CreateSize(ctx context.Context, size sizeModel.CreateSize) (sizeModel.Sizes, error) {
	return s.repo.CreateSize(ctx, size)
}

func (s *SizeService) UpdateSize(ctx context.Context, sizeID string, size sizeModel.CreateSize) (sizeModel.Sizes, error) {
	sizeFinded, err := s.repo.FindBySizeID(ctx, sizeID)
	if err != nil {
		return sizeModel.Sizes{}, err
	}
	return s.repo.UpdateSize(ctx, sizeFinded.ID.String(), size)
}

func (s *SizeService) DeleteSizeByID(ctx context.Context, sizeID string) (sizeModel.Sizes, error) {
	sizeFinded, err := s.repo.FindBySizeID(ctx, sizeID)

	if err != nil {
		return sizeModel.Sizes{}, err
	}
	return s.repo.DeleteSizeByID(ctx, sizeFinded.ID.String())
}

func (s *SizeService) GetSizeByID(ctx context.Context, sizeID string) (sizeModel.Sizes, error) {
	sizeFinded, err := s.repo.FindBySizeID(ctx, sizeID)

	if err != nil {
		return sizeModel.Sizes{}, err
	}
	return s.repo.GetSizeByID(ctx, sizeFinded.ID.String())
}
