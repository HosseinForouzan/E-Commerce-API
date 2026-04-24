package cartservice

import (
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/HosseinForouzan/E-Commerce-API/service/productservice"
)

type Repository interface {
	AddItem(req param.AddItemRequest) error
}

type Service struct {
	CartRepo Repository
	ProductRepo productservice.Repository
}

func New(CartRepo Repository, ProductRepo productservice.Repository) Service {
	return Service{CartRepo: CartRepo, ProductRepo: ProductRepo}
}

func (s Service) AddItem(req param.AddItemRequest) (error) {
	proudct, err := s.ProductRepo.GetProductByID(req.ProductID)
	if err != nil {
		return  fmt.Errorf("unexpected error: %w", err)
	}

	if proudct.Stock < req.Quantity {
		return  fmt.Errorf("not enough quantity.")
	}

	if err != nil {
		return  fmt.Errorf("unexpected error: %w", err)
	}

	return s.CartRepo.AddItem(req)

}