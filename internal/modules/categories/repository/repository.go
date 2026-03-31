package categoriesRepo

import (
	"context"
	"errors"
	categoryModel "go-minimal/internal/modules/categories/model"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type CategoriesRepoI interface {
	GetAllCategories(ctx context.Context) ([]categoryModel.Categories, error)
	CreateCategory(ctx context.Context, category categoryModel.CreateCategory) (categoryModel.Categories, error)
	FindByCategoryID(ctx context.Context, categoryID string) (categoryModel.Categories, error)
	UpdateCategory(ctx context.Context, categoryID string, category categoryModel.CreateCategory) (categoryModel.Categories, error)

	DeleteCategoryById(ctx context.Context, categoryID string) (categoryModel.Categories, error)
	GeyByCategoryID(ctx context.Context, categoryID string) (categoryModel.Categories, error)
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
func (r *CategoriesRepo) FindByCategoryID(ctx context.Context, categoryID string) (categoryModel.Categories, error) {
	var category categoryModel.Categories

	findByMaterialQuery := `SELECT id, name, is_active, slug, created_at, updated_at FROM categories WHERE id = $1`

	err := r.db.QueryRow(ctx, findByMaterialQuery, categoryID).Scan(&category.ID, &category.Name, &category.IsActive, &category.Slug, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return categoryModel.Categories{}, errors.New("requested  category did not exist")
	}
	return category, nil
}

func (r *CategoriesRepo) CreateCategory(ctx context.Context, category categoryModel.CreateCategory) (categoryModel.Categories, error) {
	var created categoryModel.Categories
	query := `INSERT INTO categories (name, is_active, slug) VALUES ($1, $2, $3) RETURNING id, name , is_active, slug, created_at, updated_at `

	err := r.db.QueryRow(ctx, query, category.Name, category.IsActive, category.Slug).Scan(&created.ID, &created.Name, &created.IsActive, &created.Slug, &created.CreatedAt, &created.UpdatedAt)
	log.Println("QUERY ERR CREATE", err)

	if err != nil {
		// Detect Postgres error
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {

			case "23505": // unique violation
				switch pgErr.ConstraintName {
				case "unique_category_name":
					return created, errors.New("category name already exists")
				case "categories_slug_key":
					return created, errors.New("this slug is already taken")
				default:
					return created, errors.New("duplicate value")
				}

			case "23502": // not null violation
				return created, errors.New("missing required field")

			case "23514": // check constraint
				return created, errors.New("invalid field value")
			}
		}
		return created, err
	}

	return created, nil

}

func (r *CategoriesRepo) UpdateCategory(ctx context.Context, categoryID string, category categoryModel.CreateCategory) (categoryModel.Categories, error) {
	var updated categoryModel.Categories

	// check if material name already exits
	var existingCategoryID int
	nameCheckQuery := `SELECT id from categories WHERE name = $1 AND id != $2`

	checkNameErr := r.db.QueryRow(ctx, nameCheckQuery, category.Name, categoryID).Scan(&existingCategoryID)

	if checkNameErr == nil {
		// row found - email belongs to someone else
		return updated, errors.New("category name already exists")
	}

	updateCategoryQuery := `UPDATE categories SET name = $1,is_active = $2, slug = $3 , updated_at = NOW() WHERE id = $4 RETURNING id, name, is_active, slug,created_at, updated_at`

	updateCategoryQueryErr := r.db.QueryRow(ctx, updateCategoryQuery, category.Name, category.IsActive, category.Slug, categoryID).Scan(&updated.ID, &updated.Name, &updated.IsActive, &updated.Slug, &updated.CreatedAt, &updated.UpdatedAt)

	if updateCategoryQueryErr != nil {
		if pgErr, ok := updateCategoryQueryErr.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505":
				switch pgErr.ConstraintName {
				case "unique_category_name":
					return updated, errors.New("category name already exists")
				case "categories_slug_key":
					return updated, errors.New("this slug is already taken")
				default:
					return updated, errors.New("duplicate value")
				}
			case "23502":
				return updated, errors.New("missing required field")
			}
		}
		return updated, updateCategoryQueryErr
	}

	return updated, nil

}

func (r *CategoriesRepo) DeleteCategoryById(ctx context.Context, categoryID string) (categoryModel.Categories, error) {

	var deletedCategory categoryModel.Categories

	deleteCategoryQuery := `DELETE FROM categories WHERE ID = $1 RETURNING id,name, is_active,slug, created_at,updated_at`

	err := r.db.QueryRow(ctx, deleteCategoryQuery, categoryID).Scan(&deletedCategory.ID, &deletedCategory.Name, &deletedCategory.IsActive, &deletedCategory.Slug, &deletedCategory.CreatedAt, &deletedCategory.UpdatedAt)

	if err != nil {
		return deletedCategory, errors.New("failed to delete category")
	}

	return deletedCategory, nil

}

func (r *CategoriesRepo) GeyByCategoryID(ctx context.Context, categoryID string) (categoryModel.Categories, error) {
	var category categoryModel.Categories

	query := `SELECT id, name, is_active,slug, created_at, updated_at FROM categories WHERE id=$1`

	err := r.db.QueryRow(ctx, query, categoryID).Scan(&category.ID, &category.Name, &category.IsActive, &category.Slug, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return categoryModel.Categories{}, errors.New("category did not found")
	}

	return category, nil
}
