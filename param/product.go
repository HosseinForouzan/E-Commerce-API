package param

import "github.com/HosseinForouzan/E-Commerce-API/entity"

type ProductRequest struct {
	CategoryID uint8
	Search     string
}

type ProductResponse struct {
	Product [] entity.Product
}

type ProductByIDRequest struct {
	ProductID uint
}

type ProductByIDResponse struct {
	Product entity.Product
}

type AddProductRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price uint `json:"price"`
	Stock uint `json:"stock"`
	CategoryID uint `json:"category_id"`

}


type AddProductResponse struct {
	Name string
}

type UpdateProductRequest struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price uint `json:"price"`
	Stock uint `json:"stock"`
	CategoryID uint `json:"category_id"`
}

type UpdateProductResponse struct {
	ID uint
	Name string
}
