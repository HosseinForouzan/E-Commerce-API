package userservice

import (
	"fmt"
	"time"

	"github.com/HosseinForouzan/E-Commerce-API/entity"
	"github.com/HosseinForouzan/E-Commerce-API/param"
)

type Repository interface {
	Register(user entity.User) (entity.User, error)
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
	fmt.Println(user)

	createdUser, err := s.repo.Register(user)
	fmt.Println(createdUser, "Salam")

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
