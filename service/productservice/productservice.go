package productservice

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/entity"
	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	GetProducts(p param.ProductRequest) ([] entity.Product, error)
	GetProductByID(id uint) (entity.Product, error)
	CreateProduct(p entity.Product) (entity.Product, error)
	CreateCategory(c entity.Category) (entity.Category, error)
	UpdateProduct(p entity.Product)(entity.Product, error)
	DeleteProduct(id uint) error
	DecreaseStockTx(ctx context.Context, tx pgx.Tx, productID uint, quantity uint) error
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

func (s Service) AddProduct(req param.AddProductRequest) (param.AddProductResponse, error) {
	product := entity.Product{
		ID: 0,
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		Stock: req.Stock,
		CategoryID: req.CategoryID,
	}

	createdProduct, err := s.repo.CreateProduct(product)
	if err != nil {
		return param.AddProductResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.AddProductResponse{
		Name: createdProduct.Name,
	}, nil


}

func(s Service) AddCategory(req param.AddCategoryRequest) (param.AddCategoryResponse, error) {
	category := entity.Category{
		Name: req.Name,
	}

	createdCategory, err := s.repo.CreateCategory(category)
	if err != nil {
		return param.AddCategoryResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.AddCategoryResponse{
		Category: createdCategory,
	}, nil
}

func (s Service) UpdateProduct(req param.UpdateProductRequest) (param.UpdateProductResponse, error) {
	prodcut := entity.Product{
		ID: req.ID,
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		Stock: req.Stock,
		CategoryID: req.CategoryID,
	}
	createdProduct, err := s.repo.UpdateProduct(prodcut)
	if err != nil {
		return param.UpdateProductResponse{}, fmt.Errorf(err.Error())
	}

	return param.UpdateProductResponse{
		ID: createdProduct.ID,
		Name: createdProduct.Name,
	}, nil
}

func (s Service) DeleteProduct(req param.ProductByIDRequest) (error) {
	if doesProductExist, _ := s.repo.GetProductByID(req.ProductID); doesProductExist.ID == 0 {
		return fmt.Errorf("this product does'nt exist")
	}

	err := s.repo.DeleteProduct(req.ProductID)
	if err != nil {
		return fmt.Errorf("unexpected error: %w", err)
	}

	return nil


}