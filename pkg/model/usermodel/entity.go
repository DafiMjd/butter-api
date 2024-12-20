package usermodel

import (
	"butter/pkg/ctype"
	"butter/pkg/model/connectionmodel"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity struct {
	ID                  uuid.UUID      `gorm:"primary_key;column:id"`
	Username            string         `gorm:"column:username"`
	Password            string         `gorm:"column:password"`
	Email               string         `gorm:"column:email"`
	Name                string         `gorm:"column:name"`
	Birthdate           ctype.NullDate `gorm:"column:birthdate"`
	CreatedAt           time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt           gorm.DeletedAt
	FollowingConnection connectionmodel.ConnectionEntity `gorm:"foreignKey:followee_id;references:id"`
	IsFollowed          bool                             `gorm:"-"`
}

func (a *UserEntity) TableName() string {
	return "butter.users"
}

func ToUserResponse(entity UserEntity) UserResponse {
	return UserResponse{
		Id:         entity.ID.String(),
		Username:   entity.Username,
		Name:       entity.Name,
		Email:      entity.Email,
		Birthdate:  entity.Birthdate,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
		IsFollowed: entity.IsFollowed,
	}
}

func ToUserResponses(users []UserEntity) []UserResponse {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
