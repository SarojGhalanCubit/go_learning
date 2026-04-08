package productService

import (
	"context"
	"errors"
	"fmt"
	categoriesRepo "go-minimal/internal/modules/categories/repository"
	materialRepository "go-minimal/internal/modules/materials/repository"
	productModel "go-minimal/internal/modules/products/model"
	productRepo "go-minimal/internal/modules/products/repository"
	"go-minimal/internal/utils"
	"log"
	"strings"
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
	_, err := s.materialRepo.FindByMaterialID(ctx, req.MaterialID.String())
	if err != nil {
		return productModel.ProductResponse{}, errors.New("invalid material: material does not exist")
	}

	_, catErr := s.categoryRepo.FindByCategoryID(ctx, req.CategoryID.String())
	if catErr != nil {
		return productModel.ProductResponse{}, errors.New("invalid category: category does not exist")
	}

	slug, err := utils.GenerateUniqueSlug(req.Name)
	if err != nil {
		return productModel.ProductResponse{}, err
	}

	req.Slug = slug

	return s.repo.CreateProduct(ctx, req)
}

func (s *ProductService) UpdateProductByID(ctx context.Context, product productModel.CreateProduct, productID string) (productModel.ProductResponse, error) {
	productFinded, err := s.repo.GetByID(ctx, productID)

	if err != nil {
		return productModel.ProductResponse{}, errors.New("invalid product: product does not exist")
	}

	_, materialErr := s.materialRepo.FindByMaterialID(ctx, product.MaterialID.String())
	if materialErr != nil {
		return productModel.ProductResponse{}, errors.New("invalid material: material does not exist")
	}

	_, catErr := s.categoryRepo.FindByCategoryID(ctx, product.CategoryID.String())
	if catErr != nil {
		return productModel.ProductResponse{}, errors.New("invalid category: category does not exist")
	}

	nameChange := product.Name != productFinded.Name
	slugIsMissing := strings.TrimSpace(productFinded.Slug) == ""

	if nameChange || slugIsMissing {
		uniqueSlug, err := utils.GenerateUniqueSlug(product.Name)

		if err != nil {
			return productModel.ProductResponse{}, fmt.Errorf("failed to generate slug: %w", err)
		}

		product.Slug = uniqueSlug
	} else {
		product.Slug = productFinded.Slug
	}
	log.Println("product slug :: ", product)

	return s.repo.UpdateProductByID(ctx, product, productFinded.ID.String())

}
