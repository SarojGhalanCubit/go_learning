package repository

import (
	"context"
	"go-minimal/internal/model"
	"github.com/jackc/pgx/v5"
)

type UserRepositoryI interface {
	GetAll() ([]model.UserResponse, error)
	Create(user model.User) (model.UserResponse, error)
}

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() ([]model.UserResponse, error) {

	rows, err := r.db.Query(context.Background(),
		"SELECT id, name, age, email, phone_number FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.UserResponse

	for rows.Next() {
		var user model.UserResponse

		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Create(user model.User) (model.UserResponse, error) {

	var created model.UserResponse

	query := `
		INSERT INTO users (name, age, email,phone_number,password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, age,email,phone_number
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Age,
		user.Email,
		user.Phone,
		user.Password,
	).Scan(&created.ID,
		&created.Name,
		&created.Age,
		&created.Email,
		&created.Phone,)

	if err != nil {
		return created, err
	}

	return created, nil
}
