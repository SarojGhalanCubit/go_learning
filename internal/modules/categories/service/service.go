package categoriesService

import (
	"context"
	categoryModel "go-minimal/internal/modules/categories/model"
	categoriesRepo "go-minimal/internal/modules/categories/repository"
)

type CategoriesService struct {
	repo categoriesRepo.CategoriesRepoI
}

func NewCategoriesService(repo categoriesRepo.CategoriesRepoI) *CategoriesService {
	return &CategoriesService{repo: repo}
}

func (s *CategoriesService) GetAllCategories(ctx context.Context) ([]categoryModel.Categories, error) {
	return s.repo.GetAllCategories(ctx)
}
