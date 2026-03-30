package authService

import (
	"errors"
	"go-minimal/internal/modules/users/model"
	"go-minimal/internal/modules/users/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo userRepository.UserRepositoryI
}

func NewAuthService(repo userRepository.UserRepositoryI) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(email, password string) (model.User, error) {

	user, err := s.repo.FindByEmail(email)

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
