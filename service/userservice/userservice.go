package userservice

import (
	"fmt"
	"time"

	"github.com/HosseinForouzan/E-Commerce-API/entity"
	"github.com/HosseinForouzan/E-Commerce-API/param"
)

type Repository interface {
	Register(user entity.User) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByID(id uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}
type Service struct {
	repo Repository
	auth AuthGenerator
}

func New(repo Repository, auth AuthGenerator) Service {
	return Service{repo: repo, auth: auth}
}

func (s Service) Register(req param.RegisterRequest) (param.RegisterRespone, error) {	

	user := entity.User{
		ID: 0,
		Name: req.Name,
		Password: req.Password,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := s.repo.Register(user)

	if err != nil {
		return param.RegisterRespone{}, fmt.Errorf("unexpected error %w", err)
	}

	return param.RegisterRespone{User: param.UserInfo{
		ID: createdUser.ID,
		Name: createdUser.Name,
		Email: createdUser.Email,
		PhoneNumber: createdUser.PhoneNumber,
	}}, nil
}

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
		user, err := s.repo.GetUserByEmail(req.Email)
		if err != nil {
			return param.LoginResponse{}, fmt.Errorf("can't retrieve user %w", err)
		}

		if req.Password != user.Password {
			return param.LoginResponse{}, fmt.Errorf("email or password is incorrect.")
		}

		accessToken, err := s.auth.CreateAccessToken(user)
		if err != nil {
			return param.LoginResponse{}, fmt.Errorf("unexpected error %w", err)
		}

		refreshToken, err := s.auth.CreateRefreshToken(user)
		if err != nil {
			return param.LoginResponse{}, fmt.Errorf("unexpted error %w", err)
		}
	

		return param.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return param.ProfileResponse{}, fmt.Errorf("can't retrieve user id %w", err)
	}

	return param.ProfileResponse{Name: user.Name}, nil
}

