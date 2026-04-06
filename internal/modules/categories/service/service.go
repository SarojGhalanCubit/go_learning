package categoriesService

import (
	"context"
	"fmt"
	categoryModel "go-minimal/internal/modules/categories/model"
	categoriesRepo "go-minimal/internal/modules/categories/repository"
	"go-minimal/internal/utils"
	"log"
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

func (s *CategoriesService) CreateCategory(ctx context.Context, category categoryModel.CreateCategory) (categoryModel.Categories, error) {
	uniqueSlug, err := utils.GenerateUniqueSlug(category.Name)
	if err != nil {
		return categoryModel.Categories{}, fmt.Errorf("failed to generate slug: %w", err)
	}

	category.Slug = uniqueSlug
	return s.repo.CreateCategory(ctx, category)
}

func (s *CategoriesService) UpdateCategory(ctx context.Context, categoryID string, category categoryModel.CreateCategory) (categoryModel.Categories, error) {
	categoryFinded, err := s.repo.FindByCategoryID(ctx, categoryID)
	if err != nil {
		return categoryModel.Categories{}, err
	}
	if category.Name != categoryFinded.Name {
		uniqueSlug, err := utils.GenerateUniqueSlug(category.Name)
		if err != nil {
			return categoryModel.Categories{}, fmt.Errorf("failed to generate slug: %w", err)
		}

		category.Slug = uniqueSlug

	} else {
		category.Slug = categoryFinded.Slug
	}
	return s.repo.UpdateCategory(ctx, categoryFinded.ID.String(), category)
}

func (s *CategoriesService) DeleteCategoryById(ctx context.Context, categoryID string) (categoryModel.Categories, error) {
	categoryFinded, err := s.repo.FindByCategoryID(ctx, categoryID)

	log.Println("DELETE ERR :: ", err)

	if err != nil {
		return categoryModel.Categories{}, err
	}
	return s.repo.DeleteCategoryById(ctx, categoryFinded.ID.String())
}

func (s *CategoriesService) GeyByCategoryID(ctx context.Context, categoryID string) (categoryModel.Categories, error) {

	materialFinded, err := s.repo.FindByCategoryID(ctx, categoryID)

	if err != nil {
		return categoryModel.Categories{}, err
	}

	return s.repo.GeyByCategoryID(ctx, materialFinded.ID.String())
}
