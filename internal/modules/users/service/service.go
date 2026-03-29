package service

import (
	"errors"
	"go-minimal/internal/modules/users/model"
	"go-minimal/internal/modules/users/repository"
	"log"
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

func (s *UserService) GetUsers() ([]model.UserResponse, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user model.User) (model.UserResponse, error) {

	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(userID int) (model.UserResponse, error) {
	user, err := s.repo.GetUserById(userID)
	if err != nil {
		return model.UserResponse{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(userID int, user model.UserResponse) (model.UserResponse, error) {
	return s.repo.UpdateUser(userID, user)
}

func (s *UserService) DeleteUser(userID int) (model.UserResponse, error) {
	user, err := s.repo.FindByUserID(userID)
	log.Println("USER IS ERR : ", err)

	if err != nil {
		return model.UserResponse{}, errors.New("user not found")
	}
	return s.repo.DeleteUser(user.ID)
}
