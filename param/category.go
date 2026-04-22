package param

import "github.com/HosseinForouzan/E-Commerce-API/entity"

type AddCategoryRequest struct {
	Name string `json:"name"`
}

type AddCategoryResponse struct {
	Category entity.Category `json:"category"`
}