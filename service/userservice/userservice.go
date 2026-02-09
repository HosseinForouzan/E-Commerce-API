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
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
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

	

		return param.LoginResponse{User: param.UserInfo{ID:user.ID, Email: user.Email, Name: user.Name}}, nil
}