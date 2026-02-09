package param

type RegisterRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterRespone struct {
	User UserInfo `json:"user"`
}