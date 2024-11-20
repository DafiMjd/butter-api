package postmodel

import (
	"butter/pkg/model/usermodel"
	"time"
)

type PostResponse struct {
	ID           string                 `json:"id"`
	UserId       string                 `json:"userId"`
	Content      string                 `json:"content"`
	CreatedAt    time.Time              `json:"createdAt"`
	UpdatedAt    time.Time              `json:"updatedAt"`
	UserResponse usermodel.UserResponse `json:"user"`
}
