package param

type PaymentRequest struct {
	UserID uint `json:"user_id"`
	OrderID uint `json:"order_id"`
}

type PaymentResponse struct {
	
}