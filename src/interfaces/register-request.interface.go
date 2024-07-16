package interfaces

type RegisterRequest struct {
	Username string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
