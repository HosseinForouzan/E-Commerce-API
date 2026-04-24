package param


type AddItemRequest struct {
	UserID uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}


type UpdateItemRequest struct {
	Quantity uint `json:"quantity"`
}

type CartItemResponse struct {
	ProductID uint   `json:"product_id"`
	Name      string `json:"name"`
	Price     uint   `json:"price"`
	Quantity  uint   `json:"quantity"`
	Subtotal  uint   `json:"subtotal"`
}

type CartResponse struct {
	Items []CartItemResponse `json:"items"`
	Total int                `json:"total"`
}