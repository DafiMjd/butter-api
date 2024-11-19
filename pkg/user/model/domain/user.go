package domain

import (
	"butter/pkg/ctype"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primary_key;column:id"`
	Username  string         `gorm:"column:username"`
	Password  string         `gorm:"column:password"`
	Email     string         `gorm:"column:email"`
	Name      string         `gorm:"column:name"`
	Birthdate ctype.NullDate `gorm:"column:birthdate"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
