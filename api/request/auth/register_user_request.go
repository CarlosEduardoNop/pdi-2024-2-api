package request

type RegisterUserRequest struct {
	User     string `json:"user" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
