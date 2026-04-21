package authorizationservice

import "fmt"

type Repository interface {
	IsUserAdmin(userID uint8) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CheckAccess(userID uint8) (bool, error) {
	isAdmin, err := s.repo.IsUserAdmin(userID)
	if err != nil {
		return false, fmt.Errorf("unexpected error: %w", err)
	}

	return isAdmin, nil
}