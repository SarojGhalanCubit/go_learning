package productRepo

import (
	"context"
	"errors"
	productModel "go-minimal/internal/modules/products/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type ProductRepoI interface {
	GetAllProducts(ctx context.Context) ([]productModel.ProductResponse, error)
	CreateProduct(ctx context.Context, product productModel.CreateProduct) (productModel.ProductResponse, error)
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
		err :=
			queryRow.Scan(&product.ID, &product.Name, &product.Slug, &product.Description, &product.Quantity, &product.IsActive, &product.CategoryName, &product.MaterialName)
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
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, name, slug`

	// Simple insert first
	err := r.db.QueryRow(ctx, query, p.Name, p.Description, p.Slug, p.Quantity, p.IsActive, p.CategoryID, p.MaterialID).Scan(&resp.ID, &resp.Name, &resp.Slug)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return resp, errors.New("product slug or name already exists")
		}
		return resp, err
	}
	return r.GetByID(ctx, resp.ID.String())
}

func (r *ProductRepo) GetByID(ctx context.Context, id string) (productModel.ProductResponse, error) {
	var p productModel.ProductResponse
	query := `
		SELECT p.id, p.name, p.slug, p.description, p.quantity, p.is_active, c.name, m.name 
		FROM products p 
		JOIN categories c ON p.category_id = c.id 
		JOIN materials m ON p.material_id = m.id 
		WHERE p.id = $1 AND p.deleted_at IS NULL`

	err := r.db.QueryRow(ctx, query, id).Scan(&p.ID, &p.Name, &p.Slug, &p.Description, &p.Quantity, &p.IsActive, &p.CategoryName, &p.MaterialName)
	if err != nil {
		return p, errors.New("product not found")
	}
	return p, nil
}
