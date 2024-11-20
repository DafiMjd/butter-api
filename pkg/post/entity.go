package post

import (
	"butter/pkg/user"
	"time"

	"gorm.io/gorm"
)

type PostEntity struct {
	ID        string    `gorm:"primary_key;column:id"`
	UserId    string    `gorm:"column:user_id"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	User      user.UserEntity `gorm:"foreignKey:user_id;references:id"`
}

func (a *PostEntity) TableName() string {
	return "butter.posts"
}
