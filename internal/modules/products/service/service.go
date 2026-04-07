package productService

import (
	"context"
	categoriesRepo "go-minimal/internal/modules/categories/repository"
	materialRepository "go-minimal/internal/modules/materials/repository"
	productModel "go-minimal/internal/modules/products/model"
	productRepo "go-minimal/internal/modules/products/repository"
	"go-minimal/internal/utils"
)

type ProductService struct {
	repo         productRepo.ProductRepoI
	categoryRepo categoriesRepo.CategoriesRepoI
	materialRepo materialRepository.MaterialRepositoryI
}

func NewProductService(
	pRepo productRepo.ProductRepoI,
	cRepo categoriesRepo.CategoriesRepoI,
	mRepo materialRepository.MaterialRepositoryI,
) *ProductService {
	return &ProductService{
		repo:         pRepo,
		categoryRepo: cRepo,
		materialRepo: mRepo,
	}
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]productModel.ProductResponse, error) {
	return s.repo.GetAllProducts(ctx)
}

func (s *ProductService) CreateProduct(ctx context.Context, req productModel.CreateProduct) (productModel.ProductResponse, error) {
	materialFinded, err := s.materialRepo.FindByMaterialID(ctx, req.MaterialID.String)
	slug, err := utils.GenerateUniqueSlug(req.Name)
	if err != nil {
		return productModel.ProductResponse{}, err
	}

	req.Slug = slug

	return s.repo.CreateProduct(ctx, req)
}
