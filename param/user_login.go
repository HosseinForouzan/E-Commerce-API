package param

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User UserInfo `json:"user"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}