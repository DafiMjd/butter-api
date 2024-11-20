package usermodel

type LoginRequest struct {
	Username string `json:"username" validate:"max=50,min=1"`
	Password string `json:"password" validate:"max=50,min=8"`
	Email    string `json:"email" validate:"max=100,min=8"`
}
