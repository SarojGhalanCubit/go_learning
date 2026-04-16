package colorsRepository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	colorModel "go-minimal/internal/modules/colors/model"
)

type ColorsRepositoryI interface {
	GetAllColors(ctx context.Context) ([]colorModel.Colors, error)
	CreateColor(ctx context.Context, color colorModel.CreateColor) (colorModel.Colors, error)
	FIndByColorID(ctx context.Context, colorID string) (colorModel.Colors, error)
	GetByColorID(ctx context.Context, colorID string) (colorModel.Colors, error)
	UpdateColor(ctx context.Context, colorID string, color colorModel.CreateColor) (colorModel.Colors, error)
	DeleteByColorID(ctx context.Context, colorID string) (colorModel.Colors, error)
}

type ColorsRepository struct {
	db *pgx.Conn
}

func NewColorsRepository(db *pgx.Conn) *ColorsRepository {
	return &ColorsRepository{db: db}
}

func (r *ColorsRepository) GetAllColors(ctx context.Context) ([]colorModel.Colors, error) {
	var colors []colorModel.Colors

	query := `SELECT id,name, hex_code, created_at FROM colors WHERE deleted_at IS NULL `

	colorsRows, err := r.db.Query(ctx, query)
	if err != nil {
		return colors, err
	}

	defer colorsRows.Close()

	for colorsRows.Next() {
		var color colorModel.Colors
		err := colorsRows.Scan(&color.ID, &color.Name, &color.HexCode, &color.CreatedAt)

		if err != nil {
			return nil, err
		}

		colors = append(colors, color)
	}

	return colors, nil
}

func (r *ColorsRepository) CreateColor(ctx context.Context, color colorModel.CreateColor) (colorModel.Colors, error) {
	var createdColor colorModel.Colors
	query := `INSERT INTO colors (name, hex_code) VALUES ($1,$2) RETURNING id, name , hex_code, created_at`

	err := r.db.QueryRow(ctx, query, color.Name, color.HexCode).Scan(&createdColor.ID, &createdColor.Name, &createdColor.HexCode, &createdColor.CreatedAt)

	if err != nil {
		// Detect Postgres error
		if pgErr, ok := err.(*pgconn.PgError); ok {

			switch pgErr.Code {

			case "23505": // unique violation
				switch pgErr.ConstraintName {
				case "unique_color_name":
					return createdColor, errors.New("color name already exists")
				default:
					return createdColor, errors.New("duplicate value")
				}

			case "23502": // not null violation
				return createdColor, errors.New("missing required field")

			case "23514": // check constraint
				return createdColor, errors.New("invalid field value")
			}
		}

		return createdColor, err
	}

	return createdColor, nil
}

func (r *ColorsRepository) FIndByColorID(ctx context.Context, colorID string) (colorModel.Colors, error) {
	var color colorModel.Colors

	query := `SELECT id, name, hex_code, created_at  FROM colors WHERE id = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(ctx, query, colorID).Scan(&color.ID, &color.Name, &color.HexCode, &color.CreatedAt)

	if err != nil {
		return colorModel.Colors{}, errors.New("requested color did not exist")
	}
	return color, nil
}

func (r *ColorsRepository) GetByColorID(ctx context.Context, colorID string) (colorModel.Colors, error) {
	var color colorModel.Colors

	query := `SELECT id, name,  hex_code,created_at  FROM colors WHERE id=$1 AND deleted_at IS NULL `

	err := r.db.QueryRow(ctx, query, colorID).Scan(&color.ID, &color.Name, &color.HexCode, &color.CreatedAt)

	if err != nil {
		return colorModel.Colors{}, errors.New("color did not found")
	}

	return color, nil
}

func (r *ColorsRepository) UpdateColor(ctx context.Context, colorID string, color colorModel.CreateColor) (colorModel.Colors, error) {

	var updated colorModel.Colors
	updateColorQuery := `UPDATE colors SET name = $1,hex_code = $2 WHERE id = $3 AND deleted_at IS NULL RETURNING id, name, hex_code, created_at `
	updateColorQueryErr := r.db.QueryRow(ctx, updateColorQuery, color.Name, color.HexCode, colorID).Scan(&updated.ID, &updated.Name, &updated.HexCode, &updated.CreatedAt)
	if updateColorQueryErr != nil {
		return updated, updateColorQueryErr

	}
	return updated, nil
}

func (r *ColorsRepository) DeleteByColorID(ctx context.Context, colorID string) (colorModel.Colors, error) {

	var deletedColors colorModel.Colors
	query := `UPDATE colors SET deleted_at = NOW() WHERE ID = $1 AND deleted_at IS NULL RETURNING id,name, hex_code, created_at `

	err := r.db.QueryRow(ctx, query, colorID).Scan(&deletedColors.ID, &deletedColors.Name, &deletedColors.HexCode, &deletedColors.CreatedAt)

	if err != nil {
		return deletedColors, errors.New("failed to delete colors")
	}

	return deletedColors, nil

}
