package service

import (
	"go-minimal/model"
	"go-minimal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	if repo == nil {
		panic("Repository cannot be nil ") 
	}
	return &UserService{repo: repo}
}


func ( s *UserService ) GetUsers() []model.User {
	return  s.repo.GetAll()
}


func (s *UserService) CreateUser(user model.User) model.User {
	return s.repo.Create(user)
}
