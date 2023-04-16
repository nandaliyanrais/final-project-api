package request

// UserRegisterRequest represents the user register request
type UserRegisterRequest struct {
	Username string `binding:"required" json:"username" form:"username"`
	Age      int    `binding:"required,gt=8" json:"age" form:"age"`
	Email    string `binding:"required,email" json:"email" form:"email"`
	Password string `binding:"required,min=6" json:"password" form:"password"`
}

// UserLoginRequest represents the user login request
type UserLoginRequest struct {
	Email    string `binding:"required" json:"email" form:"email"`
	Password string `binding:"required" json:"password" form:"password"`
}
