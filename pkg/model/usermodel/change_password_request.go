package usermodel

type ChangePasswordRequest struct {
	ID          string `json:"id"`
	OldPassword string `json:"oldPassword" validate:"max=50,min=8"`
	NewPassword string `json:"newPassword" validate:"max=50,min=8"`
}
