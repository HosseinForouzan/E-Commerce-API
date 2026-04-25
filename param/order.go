package param

type OrderResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type OrderDetailResponse struct {
	ID     int                  `json:"id"`
	Status string               `json:"status"`
	Total  int                  `json:"total"`
	Items  []OrderItemResponse  `json:"items"`
}

type OrderItemResponse struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
}