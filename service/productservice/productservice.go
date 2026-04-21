package productservice

import (
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/entity"
	"github.com/HosseinForouzan/E-Commerce-API/param"
)

type Repository interface {
	GetProducts(p param.ProductRequest) ([] entity.Product, error)
	GetProductByID(id uint8) (entity.Product, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Product(req param.ProductRequest) (param.ProductResponse, error) {
	products, err := s.repo.GetProducts(req)
	if err != nil {
		return param.ProductResponse{}, fmt.Errorf("can't get products: %w", err)
	}


	return param.ProductResponse{
		Product: products,
	}, nil

}

func (s Service) ProductByID(req param.ProductByIDRequest) (param.ProductByIDResponse, error) {
	product, err := s.repo.GetProductByID(req.ProductID)
	if err != nil {
		return param.ProductByIDResponse{}, fmt.Errorf(err.Error())
	}

	return param.ProductByIDResponse{Product: product}, nil
}