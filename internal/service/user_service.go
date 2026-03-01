package service

import (
	"go-minimal/internal/model"
	"go-minimal/internal/repository"
	"errors"
)

type UserService struct {
	repo repository.UserRepositoryI
}

func NewUserService(repo repository.UserRepositoryI) *UserService {
	if repo == nil {
		panic("repository cannot be nil")
	}

	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user model.User) (model.User, error) {
	if err := s.validateUser(user); err != nil {
		return model.User{},err
	}
	return s.repo.Create(user)
}

func (s *UserService) validateUser(user model.User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	if user.Age < 18 {
		return errors.New("user must be 18+")
	}

	return nil
}
