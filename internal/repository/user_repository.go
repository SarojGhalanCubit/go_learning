package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"go-minimal/internal/model"
	"github.com/jackc/pgx/v5"
)

type UserRepositoryI interface {
	GetAll() ([]model.UserResponse, error)
	Create(user model.User) (model.UserResponse, error)
	FindByEmail(email string) (model.User,error)
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
		INSERT INTO users (name, age, email, phone_number, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, age, email, phone_number
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Age,
		user.Email,
		user.Phone,
		user.Password,
	).Scan(
		&created.ID,
		&created.Name,
		&created.Age,
		&created.Email,
		&created.Phone,
	)

	if err != nil {

		// Detect Postgres error
		if pgErr, ok := err.(*pgconn.PgError); ok {

			switch pgErr.Code {

			case "23505": // unique violation
				switch pgErr.ConstraintName {
				case "user_email_unique":
					return created, errors.New("email already exists")
				case "user_phone_unique":
					return created, errors.New("phone already exists")
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


func (r *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User

	query := `SELECT id,name,email,password FROM users WHERE email=$1`

	err :=  r.db.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name,&user.Email,&user.Password)

	if err != nil {
		return model.User{},err
	}
	return user, nil
}

