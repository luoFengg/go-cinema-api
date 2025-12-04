package web

type UserLoginRequest struct {
	Identifier string `json:"identifier" binding:"required"` // bisa email atau username
	Password   string `json:"password" binding:"required,min=6,max=100"`
}