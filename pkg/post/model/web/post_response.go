package web

import (
	"butter/pkg/user/model/web"
	"time"
)

type PostResponse struct {
	ID           string           `json:"id"`
	UserId       string           `json:"userId"`
	Content      string           `json:"content"`
	CreatedAt    time.Time        `json:"createdAt"`
	UpdatedAt    time.Time        `json:"updatedAt"`
	UserResponse web.UserResponse `json:"user"`
}
