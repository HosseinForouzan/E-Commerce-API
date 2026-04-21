package entity

import "time"

type Product struct {
	ID          uint8
	Name        string
	Description string
	Price       int
	Stock       int
	CategoryID  uint8
	CreatedAt   time.Time
	UpdatedAt time.Time
}
