package repository

import (
	"context"
	"go-minimal/internal/model"
	"github.com/jackc/pgx/v5"
	"log"
)

type UserRepositoryI interface {
	GetAll() ([]model.User, error)
	Create(user model.User) (model.User, error)
}

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() ([]model.User, error) {

	rows, err := r.db.Query(context.Background(),
		"SELECT id, name, age FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Create(user model.User) (model.User, error) {

	query := `
		INSERT INTO users (name, age, email,phone_number,password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	log.Println("Server Running on http://localhost:8082")

	err := r.db.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Age,
		user.Email,
		user.Phone,
		user.Password,
	).Scan(&user.ID)

	if err != nil {
		return user, err
	}

	return user, nil
}
