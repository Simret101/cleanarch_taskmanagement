package Repositories

import (
	"errors"
	"task/Domain"
)

type userRepository struct {
	users []Domain.User
}

func NewUserRepository() Domain.UserRepository {
	return &userRepository{users: []Domain.User{}}
}

func (r *userRepository) GetUserByUsername(username string) (*Domain.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *userRepository) CreateUser(user *Domain.User) error {
	r.users = append(r.users, *user)
	return nil
}
