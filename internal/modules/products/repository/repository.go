package productRepo

import (
	"context"
	"errors"
	productModel "go-minimal/internal/modules/products/model"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type ProductRepoI interface {
	GetAllProducts(ctx context.Context) ([]productModel.ProductResponse, error)
	CreateProduct(ctx context.Context, product productModel.CreateProduct) (productModel.ProductResponse, error)
	GetByID(ctx context.Context, productID string) (productModel.ProductResponse, error)
	DeleteProductByID(ctx context.Context, productID string) (productModel.ProductResponse, error)
	UpdateProductByID(ctx context.Context, product productModel.CreateProduct, productID string) (productModel.ProductResponse, error)
}

type ProductRepo struct {
	db *pgx.Conn
}

func NewProductRepo(db *pgx.Conn) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) GetAllProducts(ctx context.Context) ([]productModel.ProductResponse, error) {
	query := `SELECT p.id,p.name,p.slug,p.description,p.quantity,p.is_active,c.name AS category_name,m.name AS material_name FROM products p JOIN categories c ON p.category_id = c.id JOIN materials m ON p.material_id = m.id WHERE p.deleted_at IS NULL AND p.is_active = true`

	var products []productModel.ProductResponse
	queryRow, err := r.db.Query(ctx, query)
	if err != nil {
		return products, err
	}
	for queryRow.Next() {
		var product productModel.ProductResponse
		err := queryRow.Scan(&product.ID, &product.Name, &product.Slug, &product.Description, &product.Quantity, &product.IsActive, &product.CategoryName, &product.MaterialName)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepo) CreateProduct(ctx context.Context, p productModel.CreateProduct) (productModel.ProductResponse, error) {
	var resp productModel.ProductResponse

	query := `
		INSERT INTO products (name, description,slug, quantity, is_active, category_id, material_id) 
		VALUES ($1, $2, $3, $4, $5, $6,$7) 
		RETURNING id, name, slug`

	// Simple insert first
	err := r.db.QueryRow(ctx, query, p.Name, p.Description, p.Slug, p.Quantity, p.IsActive, p.CategoryID, p.MaterialID).Scan(&resp.ID, &resp.Name, &resp.Slug)

	if err != nil {

		// Detect Postgres error
		if pgErr, ok := err.(*pgconn.PgError); ok {

			switch pgErr.Code {

			case "23505": // unique violation
				switch pgErr.ConstraintName {
				case "unique_product_name":
					return resp, errors.New("product name already exists")
				default:
					return resp, errors.New("duplicate value")
				}

			case "23502": // not null violation
				return resp, errors.New("missing required field")

			case "23514": // check constraint
				return resp, errors.New("invalid field value")
			}
		}

		return resp, err
	}
	return r.GetByID(ctx, resp.ID.String())
}

func (r *ProductRepo) UpdateProductByID(ctx context.Context, product productModel.CreateProduct, productID string) (productModel.ProductResponse, error) {

	var updated productModel.ProductResponse
	_, getByIDerr := r.GetByID(ctx, productID)
	if getByIDerr != nil {
		return updated, getByIDerr
	}

	query := `UPDATE products SET name = $1, description = $2, quantity = $3,slug = $4,is_active = $5,material_id = $6,category_id = $7,updated_at = NOW() WHERE id = $8 RETURNING id,name,slug`

	updateProductQueryErr := r.db.QueryRow(ctx, query, product.Name, product.Description, product.Quantity, product.Slug, product.IsActive, product.MaterialID, product.CategoryID, productID).Scan(&updated.ID, &updated.Name, &updated.Slug)

	if updateProductQueryErr != nil {
		if pgErr, ok := updateProductQueryErr.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505":
				switch pgErr.ConstraintName {
				case "unique_product_name":
					return updated, errors.New("product name already exists")
				default:
					return updated, errors.New("duplicate value")
				}
			case "23502":
				return updated, errors.New("missing required field")
			}
		}
		return updated, updateProductQueryErr
	}

	return r.GetByID(ctx, updated.ID.String())

}

func (r *ProductRepo) GetByID(ctx context.Context, productID string) (productModel.ProductResponse, error) {
	var p productModel.ProductResponse
	query := `
		SELECT p.id, p.name, p.slug, p.description, p.quantity, p.is_active, c.name, m.name 
		FROM products p 
		JOIN categories c ON p.category_id = c.id 
		JOIN materials m ON p.material_id = m.id 
		WHERE p.id = $1 AND p.deleted_at IS NULL`

	err := r.db.QueryRow(ctx, query, productID).Scan(&p.ID, &p.Name, &p.Slug, &p.Description, &p.Quantity, &p.IsActive, &p.CategoryName, &p.MaterialName)
	if err != nil {
		return p, errors.New("product not found")
	}
	return p, nil
}

func (r *ProductRepo) DeleteProductByID(ctx context.Context, productID string) (productModel.ProductResponse, error) {
	var deletedProduct productModel.ProductResponse

	query := `UPDATE products SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL RETURNING id,name`

	err := r.db.QueryRow(ctx, query, productID).Scan(&deletedProduct.ID, &deletedProduct.Name)
	if err != nil {
		return deletedProduct, errors.New("failed to delete product")
	}

	returnQuery := `
		SELECT p.id, p.name, p.slug, p.description, p.quantity, p.is_active, c.name, m.name 
		FROM products p 
		JOIN categories c ON p.category_id = c.id 
		JOIN materials m ON p.material_id = m.id 
		WHERE p.id = $1`

	returnQueryErr := r.db.QueryRow(ctx, returnQuery, productID).Scan(&deletedProduct.ID, &deletedProduct.Name, &deletedProduct.Slug, &deletedProduct.Description, &deletedProduct.Quantity, &deletedProduct.IsActive, &deletedProduct.CategoryName, &deletedProduct.MaterialName)
	log.Println("second err :: ", returnQueryErr)
	if returnQueryErr != nil {
		return deletedProduct, errors.New("product not found")
	}

	return deletedProduct, nil
}
