package postmodel

import (
	"butter/pkg/model/usermodel"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostEntity struct {
	ID        uuid.UUID `gorm:"primary_key;column:id"`
	UserId    string    `gorm:"column:user_id"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	User      usermodel.UserEntity `gorm:"foreignKey:user_id;references:id"`
}

func (a *PostEntity) TableName() string {
	return "butter.posts"
}

func ToPostResponse(post PostEntity) PostResponse {
	return PostResponse{
		ID:           post.ID.String(),
		UserId:       post.UserId,
		Content:      post.Content,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
		UserResponse: usermodel.ToUserResponse(post.User),
	}
}

func ToPostResponses(posts []PostEntity) []PostResponse {
	var userResponses []PostResponse
	for _, post := range posts {
		userResponses = append(userResponses, ToPostResponse(post))
	}

	return userResponses
}
