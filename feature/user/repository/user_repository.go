package repository

import (
	"butter/feature/user/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, user domain.User) (domain.User, error)
	Update(db *gorm.DB, user domain.User) (domain.User, error)
	Delete(db *gorm.DB, user domain.User) error
	FindById(db *gorm.DB, id string) (domain.User, error)
	FindAll(db *gorm.DB) ([]domain.User, error)
	FindBy(db *gorm.DB, query string, value interface{}) (domain.User, error)
	ChangePassword(db *gorm.DB, id string, password string) error
}
