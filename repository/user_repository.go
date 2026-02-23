package repository

import "go-minimal/model"

type UserRepository interface {
	GetAll() []model.User
	Create(user model.User) model.User
}

type InMemoryUserRepository struct {
	users []model.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []model.User{},
	}
}

func (r *InMemoryUserRepository) GetAll() []model.User {
	return r.users
}

func (r *InMemoryUserRepository) Create(user model.User) model.User {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return user
}
