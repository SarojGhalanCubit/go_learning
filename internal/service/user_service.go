package service

import (
	"go-minimal/internal/model"
	"go-minimal/internal/repository"
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
	return s.repo.Create(user)
}
