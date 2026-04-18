package entity

import "time"

type User struct {
	ID          uint
	Name        string
	Password string
	Email       string
	PhoneNumber string
	Role Role
	CreatedAt   time.Time
	UpdatedAt time.Time
}