package sizeRepository

import (
	"context"
	"errors"
	sizeModel "go-minimal/internal/modules/sizes/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type SizeRepositoryI interface {
	GetAllSizes(ctx context.Context) ([]sizeModel.Sizes, error)
	CreateSize(ctx context.Context, size sizeModel.CreateSize) (sizeModel.Sizes, error)
	UpdateSize(ctx context.Context, sizeID string, size sizeModel.CreateSize) (sizeModel.Sizes, error)
	FindBySizeID(ctx context.Context, sizeID string) (sizeModel.Sizes, error)
	DeleteSizeByID(ctx context.Context, sizeID string) (sizeModel.Sizes, error)
	GetSizeByID(ctx context.Context, sizeID string) (sizeModel.Sizes, error)
}

type SizeRepository struct {
	db *pgx.Conn
}

func NewSizeRepository(db *pgx.Conn) *SizeRepository {
	return &SizeRepository{
		db: db,
	}
}

func (r *SizeRepository) GetAllSizes(ctx context.Context) ([]sizeModel.Sizes, error) {
	var sizes []sizeModel.Sizes
	query := `SELECT id, name, sort_order,created_at FROM sizes WHERE deleted_at IS NULL`

	queryRows, err := r.db.Query(ctx, query)
	if err != nil {
		return sizes, err
	}

	defer queryRows.Close()

	for queryRows.Next() {
		var size sizeModel.Sizes

		err := queryRows.Scan(&size.ID, &size.Name, &size.SortOrder, &size.CreatedAt)

		if err != nil {
			return nil, err
		}
		sizes = append(sizes, size)
	}
	return sizes, nil
}

func (r *SizeRepository) CreateSize(ctx context.Context, size sizeModel.CreateSize) (sizeModel.Sizes, error) {

	var createdSize sizeModel.Sizes

	query := `INSERT INTO sizes (name, sort_order) VALUES ($1, $2) RETURNING id, name, sort_order, created_at `

	err := r.db.QueryRow(ctx, query, size.Name, size.SortOrder).Scan(&createdSize.ID, &createdSize.Name, &createdSize.SortOrder, &createdSize.CreatedAt)

	if err != nil {

		// Detect Postgres error
		if pgErr, ok := err.(*pgconn.PgError); ok {

			switch pgErr.Code {

			case "23505": // unique violation
				switch pgErr.ConstraintName {
				case "unique_size_name":
					return createdSize, errors.New("size name already exists")
				default:
					return createdSize, errors.New("duplicate value")
				}

			case "23502": // not null violation
				return createdSize, errors.New("missing required field")

			case "23514": // check constraint
				return createdSize, errors.New("invalid field value")
			}
		}

		return createdSize, err
	}

	return createdSize, nil
}

func (r *SizeRepository) UpdateSize(ctx context.Context, sizeID string, size sizeModel.CreateSize) (sizeModel.Sizes, error) {
	var updated sizeModel.Sizes

	query := `UPDATE sizes SET name = $1,sort_order = $2 WHERE id = $3 AND deleted_at is NULL RETURNING id, name, sort_order, created_at `

	queryErr := r.db.QueryRow(ctx, query, size.Name, size.SortOrder, sizeID).Scan(&updated.ID, &updated.Name, &updated.SortOrder, &updated.CreatedAt)

	if queryErr != nil {
		return updated, queryErr
	}

	return updated, nil

}

func (r *SizeRepository) FindBySizeID(ctx context.Context, sizeID string) (sizeModel.Sizes, error) {
	var size sizeModel.Sizes

	query := `SELECT id, name, sort_order, created_at FROM sizes WHERE id = $1 AND deleted_at IS NULL `

	err := r.db.QueryRow(ctx, query, sizeID).Scan(&size.ID, &size.Name, &size.SortOrder, &size.CreatedAt)

	if err != nil {
		return sizeModel.Sizes{}, errors.New("requested size did not exist")
	}
	return size, nil
}

func (r *SizeRepository) DeleteSizeByID(ctx context.Context, sizeID string) (sizeModel.Sizes, error) {

	var deletedSize sizeModel.Sizes

	query := `UPDATE sizes SET deleted_at = NOW() WHERE ID = $1 RETURNING id,name, sort_order, created_at`

	err := r.db.QueryRow(ctx, query, sizeID).Scan(&deletedSize.ID, &deletedSize.Name, &deletedSize.SortOrder, &deletedSize.CreatedAt)

	if err != nil {
		return deletedSize, errors.New("failed to delete size")
	}

	return deletedSize, nil

}

func (r *SizeRepository) GetSizeByID(ctx context.Context, sizeID string) (sizeModel.Sizes, error) {
	var size sizeModel.Sizes

	query := `SELECT id, name, sort_order, created_at  FROM sizes WHERE id=$1 AND deleted_at IS NULL`

	err := r.db.QueryRow(ctx, query, sizeID).Scan(&size.ID, &size.Name, &size.SortOrder, &size.CreatedAt)

	if err != nil {
		return sizeModel.Sizes{}, errors.New("size did not found")
	}

	return size, nil
}
