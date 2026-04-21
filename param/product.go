package param

import "github.com/HosseinForouzan/E-Commerce-API/entity"

type ProductRequest struct {
	CategoryID uint8
	Search     string
}

type ProductResponse struct {
	Product [] entity.Product
}