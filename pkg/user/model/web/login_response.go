package web

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	UserResponse `json:"user,omitempty"`
}
