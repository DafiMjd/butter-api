package usermodel

import (
	"butter/pkg/ctype"
	"time"
)

type UserResponse struct {
	Id        string         `json:"id"`
	Username  string         `json:"username"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Birthdate ctype.NullDate `json:"birthdate,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
