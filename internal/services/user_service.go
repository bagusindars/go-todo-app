package services

import (
	"errors"
	"simple-todo-app/internal/helpers"
	"simple-todo-app/internal/models"
	"simple-todo-app/internal/repositories"
	"time"
)

type UserService interface {
	Register(data models.UserRegisterRequest) (map[string]any, error)
	Login(data models.UserLoginRequest) (map[string]any, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) FindUserByEmail(email string) (models.Users, error) {
	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return user, errors.New("error find user by email : " + err.Error())
	}

	return user, nil
}

func (s *userService) Login(data models.UserLoginRequest) (map[string]any, error) {
	var responseData map[string]any

	if len(data.Email) == 0 {
		return responseData, errors.New("email cannot be empty")
	}

	if len(data.Password) == 0 {
		return responseData, errors.New("password cannot be empty")
	}

	user, err := s.FindUserByEmail(data.Email)

	if err != nil {
		return responseData, err
	}

	if user.Id == 0 {
		return responseData, errors.New("user not exists")
	}

	if err := helpers.CheckHashPassword(user.Password, data.Password); err != nil {
		return responseData, errors.New("email or password is incorrect")
	}

	accessToken, err := helpers.GenerateToken(uint(user.Id), user.Email, user.Name)

	if err != nil {
		return responseData, err
	}

	responseData = map[string]any{
		"user": user,
		"token": map[string]any{
			"access_token": accessToken,
		},
	}

	return responseData, nil
}

func (s *userService) Register(data models.UserRegisterRequest) (map[string]any, error) {
	var responseData map[string]any

	if len(data.Name) == 0 {
		return responseData, errors.New("name cannot be empty")
	}

	if len(data.Email) == 0 {
		return responseData, errors.New("email cannot be empty")
	}

	if len(data.Password) == 0 {
		return responseData, errors.New("password cannot be empty")
	}

	if len(data.PasswordConfirmed) == 0 {
		return responseData, errors.New("password confirmation cannot be empty")
	}

	if data.Password != data.PasswordConfirmed {
		return responseData, errors.New("password confirmation not valid")
	}

	user, err := s.FindUserByEmail(data.Email)

	if err != nil {
		return responseData, err
	}

	if user.Id != 0 {
		return responseData, errors.New("user already exists")
	}

	pass, err := helpers.MakeHash(data.Password)

	if err != nil {
		return responseData, err
	}

	user = models.Users{
		Name:      data.Name,
		Email:     data.Email,
		Password:  pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	lastId, err := s.repo.Create(user)
	if err != nil {
		return responseData, errors.New("Error register user : " + err.Error())
	}

	accessToken, err := helpers.GenerateToken(uint(lastId), user.Email, user.Name)

	if err != nil {
		return responseData, err
	}

	user.Id = lastId

	responseData = map[string]any{
		"user": user,
		"token": map[string]any{
			"access_token": accessToken,
		},
	}

	return responseData, nil
}
