package service

import (
	"context"
	"go-minimal/internal/modules/materials/model"
	"go-minimal/internal/modules/materials/repository"
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

func (s *MaterialsService) UpdateMaterial(ctx context.Context, materialID string, material model.CreateMaterial) (model.Material, error) {
	materialFinded, err := s.repo.FindByMaterialID(ctx, materialID)
	if err != nil {
		return model.Material{}, err
	}
	return s.repo.UpdateMaterial(ctx, materialFinded.ID.String(), material)
}

func (s *MaterialsService) DeleteMaterialById(ctx context.Context, materialID string) (model.Material, error) {
	materialFinded, err := s.repo.FindByMaterialID(ctx, materialID)

	if err != nil {
		return model.Material{}, err
	}
	return s.repo.DeleteMaterialById(ctx, materialFinded.ID.String())
}

func (s *MaterialsService) GeyByMaterialID(ctx context.Context, materialID string) (model.Material, error) {

	materialFinded, err := s.repo.FindByMaterialID(ctx, materialID)

	if err != nil {
		return model.Material{}, err
	}

	return s.repo.GeyByMaterialID(ctx, materialFinded.ID.String())
}
