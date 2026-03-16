package service

import (
	"errors"
	"go-minimal/internal/model"
	"go-minimal/internal/repository"

	"golang.org/x/crypto/bcrypt"
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


func (s *UserService) Login(email, password string) (model.User,error) {
	user,err := s.repo.FindByEmail(email)
	if err != nil {
		return model.User{}, errors.New("Invalid Credentials")
	}

	// Compare hashed Credentials
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil {
		return model.User{},errors.New("Invalid Credentials")
	}

	return user, nil
}

func (s *UserService) GetUserByID(userID int) (model.UserResponse, error) {
	user,err:= s.repo.GetUserById(userID)
    if err != nil {
        return model.UserResponse{}, err
    }
    return user, nil
}


func (s *UserService) UpdateUser(userID int,user model.UserResponse) (model.UserResponse, error) {
	return s.repo.UpdateUser(userID, user)
}



func (s *UserService) DeleteUser(userID int) (model.UserResponse, error) {
	return s.repo.DeleteUser(userID)
}
