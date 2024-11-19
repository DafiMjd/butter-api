package domain

import (
	"butter/pkg/ctype"
	"butter/pkg/user/model/domain"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        string `gorm:"primary_key;column:id"`
	UserId    string `gorm:"column:user_id"`
	Content   string
	Birthdate ctype.NullDate `gorm:"column:birthdate"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	User      domain.User `gorm:"foreignKey:user_id;references:id"`
}
