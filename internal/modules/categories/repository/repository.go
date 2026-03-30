package categoriesRepo

import (
	"context"
	categoryModel "go-minimal/internal/modules/categories/model"

	"github.com/jackc/pgx/v5"
)

type CategoriesRepoI interface {
	GetAllCategories(ctx context.Context) ([]categoryModel.Categories, error)
}

type CategoriesRepo struct {
	db *pgx.Conn
}

func NewCategoriesRepo(db *pgx.Conn) *CategoriesRepo {
	return &CategoriesRepo{
		db: db,
	}
}

func (r *CategoriesRepo) GetAllCategories(ctx context.Context) ([]categoryModel.Categories, error) {
	var categories []categoryModel.Categories
	query := `SELECT id, name, is_active, slug,created_at, updated_at FROM categories`

	queryRows, err := r.db.Query(ctx, query)
	if err != nil {
		return categories, err
	}

	defer queryRows.Close()

	for queryRows.Next() {
		var category categoryModel.Categories

		err := queryRows.Scan(&category.ID, &category.Name, &category.IsActive, &category.Slug, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
