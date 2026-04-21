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
	ProductID uint8
}

type ProductByIDResponse struct {
	Product entity.Product
}

