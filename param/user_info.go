package param

type UserInfo struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}