package usecase

import (
	"errors"

	dtos "github.com/savanyv/Golang-Findest/internal/dto"
	"github.com/savanyv/Golang-Findest/internal/helpers"
	"github.com/savanyv/Golang-Findest/internal/models"
	"github.com/savanyv/Golang-Findest/internal/repository"
)

type UserUsecase interface {
	Register(req *dtos.RegisterRequest) (*dtos.AuthResponse, error)
	Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) Register(req *dtos.RegisterRequest) (*dtos.AuthResponse, error) {
	// check if user already exists
	_, err := u.repo.GetUserByEmail(req.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	// hash password
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// create user
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	user, err = u.repo.CreateUser(user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	// Response
	response := &dtos.AuthResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
	}

	return response, nil
}

func (u *userUsecase) Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error) {
	// check if user exists
	user, err := u.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// compare password
	if err := helpers.ComparePassword(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate Token
	token, err := helpers.NewJWTService().GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// Response
	response := &dtos.AuthResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Token:    token,
	}

	return response, nil

}