package web

type UserCreateRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30,alphanum"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}