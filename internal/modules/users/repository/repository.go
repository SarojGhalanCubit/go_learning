package repository

import (
	"context"
	"errors"
	"go-minimal/internal/modules/users/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

/*
The interface - This defines contract of any Respository

--> It allows for dependecy injection
--> When writing unit test, we can swap the real database for a "mock" one without changing your business logic
*/
type UserRepositoryI interface {
	GetAll() ([]model.UserResponse, error)
	Create(user model.User) (model.UserResponse, error)
	FindByEmail(email string) (model.User, error)
	FindByUserID(userID int) (model.UserResponse, error)
	GetUserById(userID int) (model.UserResponse, error)
	UpdateUser(userID int, user model.UserResponse) (model.UserResponse, error)
	DeleteUser(userID int) (model.UserResponse, error)
}

/*  The Struct ---> the concreate implementation  */
type UserRepository struct {
	/* It holds pointer to pgx connection */
	/* It stores the tools ( the database connection ) neeeded to talk to Postgres */
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	/*
		It initilizes the repository with an active datgabase connection
			--> We can call this in main.go and pass the resulting repository to your service layer
	*/
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() ([]model.UserResponse, error) {

	rows, err := r.db.Query(context.Background(),
		"SELECT id, name, age, email, phone_number, role_id FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.UserResponse

	for rows.Next() {
		var user model.UserResponse

		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Phone, &user.RoleID)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Create(user model.User) (model.UserResponse, error) {

	var created model.UserResponse

	query := `
		INSERT INTO users (name, age, email, phone_number, password,role_id)
		VALUES ($1, $2, $3, $4, $5,$6)
		RETURNING id, name, age, email, phone_number, role_id
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Age,
		user.Email,
		user.Phone,
		user.Password,
		user.RoleID,
	).Scan(
		&created.ID,
		&created.Name,
		&created.Age,
		&created.Email,
		&created.Phone,
		&created.RoleID,
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

	query := `SELECT id,name,email,password,role_id FROM users WHERE email=$1`

	err := r.db.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByUserID(userID int) (model.UserResponse, error) {
	var user model.UserResponse
	query := `SELECT id,name,email,password,role_id FROM users WHERE id=$1`

	err := r.db.QueryRow(context.Background(), query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.RoleID)

	if err != nil {
		return model.UserResponse{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserById(userID int) (model.UserResponse, error) {
	var user model.UserResponse

	query := `SELECT id,name,age,email,phone_number,role_id FROM users WHERE id=$1`

	err := r.db.QueryRow(context.Background(), query, userID).Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Phone, &user.RoleID)

	if err != nil {
		return model.UserResponse{}, errors.New("User not found")
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(userID int, user model.UserResponse) (model.UserResponse, error) {

	var updated model.UserResponse

	// Check if email is already exits
	var exitingEmailID int
	emailCheckQuery := `SELECT id from users WHERE email = $1 AND id != $2`
	checkEmailErr := r.db.QueryRow(context.Background(), emailCheckQuery, user.Email, userID).Scan(&exitingEmailID)

	if checkEmailErr == nil {
		// row found - email belongs to someone else
		return updated, errors.New("email already exists")
	}

	// Check if phone number already exits
	var existingPhoneId int
	phoneCheckQuery := `SELECT id from users where phone_number = $1 AND id != $2`
	checkPhoneErr := r.db.QueryRow(context.Background(), phoneCheckQuery, user.Phone, userID).Scan(&existingPhoneId)

	if checkPhoneErr == nil {
		return updated, errors.New("phone already exists")
	}

	query := `
		UPDATE users 
		SET name = $1, age = $2, email = $3, phone_number = $4,role_id = $5
		WHERE id = $6
		RETURNING id,name,age,email,phone_number,role_id
	`
	updateUserQueryErr := r.db.QueryRow(context.Background(), query, user.Name, user.Age, user.Email, user.Phone, user.RoleID, userID).Scan(&updated.ID, &updated.Name, &updated.Age, &updated.Email, &updated.Phone, &updated.RoleID)

	if updateUserQueryErr != nil {
		if pgErr, ok := updateUserQueryErr.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505":
				switch pgErr.ConstraintName {
				case "user_email_unique":
					return updated, errors.New("email already exists")
				case "user_phone_unique":
					return updated, errors.New("phone already exists")
				default:
					return updated, errors.New("duplicate value")
				}
			case "23502":
				return updated, errors.New("missing required field")
			}
		}
		return updated, updateUserQueryErr
	}

	return updated, nil
}

func (r *UserRepository) DeleteUser(userID int) (model.UserResponse, error) {

	var deletedUser model.UserResponse

	deleteUserQuery := ` DELETE FROM users WHERE id = $1 RETURNING id, name, age, email, phone_number, role_id
    `
	deleteUserQueryErr := r.db.QueryRow(context.Background(), deleteUserQuery, userID).Scan(&deletedUser.ID, &deletedUser.Name,
		&deletedUser.Age,
		&deletedUser.Email,
		&deletedUser.Phone,
		&deletedUser.RoleID,
	)
	if deleteUserQueryErr != nil {
		return deletedUser, errors.New("failed to delete user")
	}
	return deletedUser, nil
}
