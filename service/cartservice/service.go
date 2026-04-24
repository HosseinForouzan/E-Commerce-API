package cartservice

import (
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/HosseinForouzan/E-Commerce-API/service/productservice"
)

type Repository interface {
	AddItem(req param.AddItemRequest) error
	GetCart(userID uint) (param.CartResponse, error)
	UpdateItem(req param.UpdateItemRequest) error
	DeleteItem(req param.DeleteItemRequest) error
	Clear(userID uint) error
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

func (s Service) GetCart(req param.CartRequest) (param.CartResponse, error) {
	rows, err := s.CartRepo.GetCart(req.UserID)
	if err != nil {
		return param.CartResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	resp := param.CartResponse{}
	total := 0

	for _, row := range rows.Items {
		subTotal := row.Price * row.Quantity
		total += int(subTotal)

		resp.Items = append(resp.Items, param.CartItemResponse{
			ProductID: row.ProductID,
			Name: row.Name,
			Price: row.Price,
			Quantity: row.Quantity,
			Subtotal: subTotal,
		})
	}
	resp.Total = uint(total)

	return resp, nil
}

func (s Service) UpdateItem(req param.UpdateItemRequest) error {
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

	return s.CartRepo.UpdateItem(req)
}

func (s Service) DeleteItem(req param.DeleteItemRequest) error {
	
	return s.CartRepo.DeleteItem(req)
}

func (s Service) Clear(userID uint) error {

	return s.CartRepo.Clear(userID)
}