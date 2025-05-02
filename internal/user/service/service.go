package service

import (
	"fmt"
	"learn/rest-api/internal/user/repository"
	"regexp"
)

type UserRepository interface {
	GetUserByID(id int) (repository.UserModel, error)
	CreateUser(repository.UserModel) (id int, err error)
}

type UserServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetUserByID(id int) (UserDTO, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return UserDTO{}, fmt.Errorf("faild get user %w", err)
	}
	return UserDTO{ID: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *UserServiceImpl) RegistrationUser(user RegistateUserInputDTO) (RegistateUserOutputDTO, error) {
	// Validate email format
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailRegex, user.Email)
	if err != nil {
		return RegistateUserOutputDTO{}, fmt.Errorf("failed to validate email: %w", err)
	}
	if !matched {
		return RegistateUserOutputDTO{}, ErrEmailFormat
	}

	// Validate password length
	if len(user.Password) < 8 {
		err := ErrPassowordFormat
		return RegistateUserOutputDTO{}, fmt.Errorf("password must be at least 8 characters long : %w", err)
	}

	mUser := repository.UserModel{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	id, err := s.repo.CreateUser(mUser)

	if err != nil {
		return RegistateUserOutputDTO{}, fmt.Errorf("failed create user: %w", err)
	}

	return RegistateUserOutputDTO{ID: id}, nil
}
