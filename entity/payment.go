package entity

type Payment struct {
	ID            uint
	OrderID       uint
	Amount        uint
	Status        string
	Provider      string
	TransactionID *string
}