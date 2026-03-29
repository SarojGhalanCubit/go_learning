package userLoginService

import (
	"errors"
	"go-minimal/internal/modules/users/model"
	"go-minimal/internal/modules/users/repository"

	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (model.User, error) {

	user, err := repository.UserRepositoryI.FindByEmail(_, email)

	if err != nil {
		return model.User{}, errors.New("invalid credentials")
	}

	// Compare hashed Credentials
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, errors.New("invalid credentials")
	}

	return user, nil
}
