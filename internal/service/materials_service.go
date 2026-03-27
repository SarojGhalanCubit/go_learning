package service

import (
	"context"
	"go-minimal/internal/model"
	"go-minimal/internal/repository"
)

type MaterialsService struct {
	repo repository.MaterialRepositoryI
}

func NewMaterialService(repo repository.MaterialRepositoryI) *MaterialsService {
	if repo == nil {
		panic("material repository cannot be nil")
	}
	return &MaterialsService{
		repo: repo,
	}
}

func (s *MaterialsService) GetAllMaterial(ctx context.Context) ([]model.Material, error) {
	return s.repo.GetAllMaterial(ctx)
}

func (s *MaterialsService) CreateMaterial(ctx context.Context, material model.CreateMaterial) (model.Material, error) {
	return s.repo.CreateMaterial(ctx, material)
}

func (s *MaterialsService) UpdateMaterial(ctx context.Context, materialID int, material model.CreateMaterial) (model.Material, error) {
	return s.repo.UpdateMaterial(ctx, materialID, material)
}
