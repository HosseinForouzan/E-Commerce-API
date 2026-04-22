package entity

import "time"

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       uint
	Stock       uint
	CategoryID  uint
	CreatedAt   time.Time
	UpdatedAt time.Time
}
