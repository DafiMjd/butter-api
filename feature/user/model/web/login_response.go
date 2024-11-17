package web

type LoginResponse struct {
	Token        string `json:"token"`
	UserResponse `json:"user"`
}
