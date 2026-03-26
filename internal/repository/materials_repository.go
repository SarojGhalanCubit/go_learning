package repository

import (
	"context"
	"errors"
	"go-minimal/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

/*
The interface - This defines contract of any Respository

--> It allows for dependecy injection
--> When writing unit test, we can swap the real database for a "mock" one without changing your business logic
*/
type MaterialRepositoryI interface {
	GetAllMaterial(ctx context.Context) ([]model.Material, error)
	CreateMaterial(maerial model.CreateMaterial) (model.Material, error)
}

/*  The Struct ---> the concreate implementation  */
type MaterialRepository struct {
	/* It holds pointer to pgx connection */
	/* It stores the tools ( the database connection ) neeeded to talk to Postgres */
	db *pgx.Conn
}

/*
The Constructor ---> This is Factory Function
*/
func NewMaterialRepository(db *pgx.Conn) *MaterialRepository {
	/*
		It initilizes the repository with an active datgabase connection
			--> We can call this in main.go and pass the resulting repository to your service layer
	*/
	return &MaterialRepository{
		db: db,
	}
}

func (r *MaterialRepository) GetAllMaterial(ctx context.Context) ([]model.Material, error) {
	query := `SELECT id, name, is_active,created_at, updated_at FROM materials`

	materialsRows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer materialsRows.Close()

	var materials []model.Material

	for materialsRows.Next() {
		var material model.Material

		err := materialsRows.Scan(&material.ID, &material.Name, &material.IsActive, &material.CreatedAt, &material.UpdatedAt)

		if err != nil {
			return nil, err
		}

		materials = append(materials, material)
	}

	return materials, nil
}

func (r *MaterialRepository) CreateMaterial(material model.CreateMaterial) (model.Material, error) {

	var createdMaterial model.Material

	query := `INSERT INTO materials (name, is_active) VALUES ($1,$2) RETURNING id, name , is_active, created_at, updated_at `

	err := r.db.QueryRow(context.Background(), query, material.Name, material.IsActive).Scan(&createdMaterial.ID, &createdMaterial.Name, &createdMaterial.IsActive, &createdMaterial.CreatedAt, createdMaterial.UpdatedAt)

	if err != nil {

		// Detect Postgres error
		if pgErr, ok := err.(*pgconn.PgError); ok {

			switch pgErr.Code {

			case "23505": // unique violation
				switch pgErr.ConstraintName {
				case "user_email_unique":
					return createdMaterial, errors.New("email already exists")
				case "user_phone_unique":
					return createdMaterial, errors.New("phone already exists")
				default:
					return createdMaterial, errors.New("duplicate value")
				}

			case "23502": // not null violation
				return createdMaterial, errors.New("missing required field")

			case "23514": // check constraint
				return createdMaterial, errors.New("invalid field value")
			}
		}

		return createdMaterial, err
	}

	return createdMaterial, nil

}
