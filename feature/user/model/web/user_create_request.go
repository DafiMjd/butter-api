package web

import "butter/ctype"

type UserCreateRequest struct {
	Username  string         `json:"userName" validate:"required,max=50,min=1"`
	Password  string         `json:"password" validate:"required,max=50,min=8"`
	Email     string         `json:"email" validate:"required,max=100,min=8"`
	Name      string         `json:"name" validate:"required,max=50,min=1"`
	Birthdate ctype.NullDate `json:"birthDate"`
}
