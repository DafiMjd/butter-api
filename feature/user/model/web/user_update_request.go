package web

import (
	"butter/ctype"
)

type UserUpdateRequest struct {
	ID        string         `json:"id" validate:"required"`
	Username  string         `json:"username" validate:"required,max=50,min=1"`
	Name      string         `json:"name" validate:"required,max=50,min=1"`
	Birthdate ctype.NullDate `json:"birthDate"`
}
