package entity

import "time"

type Order struct {
	ID          uint
	UserID      uint
	Status      string
	TotalAmount uint
	CreatedAt   time.Time
}

type OrderItem struct {
	ID          uint
	OrderID     uint
	ProductID   uint
	ProductName string
	UnitPrice   uint
	Quantity    uint
	Subtotal    uint
}