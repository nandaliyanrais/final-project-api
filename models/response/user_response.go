package response

// UserRegisterResponse represents the user register response
type UserRegisterResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

// UserLoginResponse represents the user login response
type UserLoginResponse struct {
	Token string `json:"token"`
}
