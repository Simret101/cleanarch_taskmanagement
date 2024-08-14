package Usecases

import (
	"errors"

	"task/Domain"
	"task/Infrastructure"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

type UserUseCase struct {
	UserRepo        Domain.UserRepository
	JWTService      Infrastructure.JWTService
	PasswordService Infrastructure.PasswordService
}

func (m *MockUserUseCase) Register(user *Domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUseCase) Login(credentials *Domain.Credentials) (string, error) {
	args := m.Called(credentials)
	return args.String(0), args.Error(1)
}

func (uc *UserUseCase) Register(user *Domain.User) error {

	existingUser, _ := uc.UserRepo.GetUserByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := uc.PasswordService.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return uc.UserRepo.CreateUser(user)
}

func (uc *UserUseCase) Login(credentials *Domain.Credentials) (string, error) {

	user, err := uc.UserRepo.GetUserByUsername(credentials.Username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !uc.PasswordService.ComparePasswords(user.Password, credentials.Password) {
		return "", errors.New("invalid credentials")
	}

	return uc.JWTService.GenerateJWT(user.Username)
}
