package repository

import (
	"butter/pkg/user/model/domain"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

// Create implements UserRepository.
func (u *UserRepositoryImpl) Create(db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.Create(&user).Error

	return user, err
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(db *gorm.DB, user domain.User) error {
	err := db.Delete(&user).Error

	return err
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll(db *gorm.DB) ([]domain.User, error) {
	var users []domain.User
	err := db.Find(&users).Error

	return users, err
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(db *gorm.DB, id string) (domain.User, error) {
	var user domain.User
	err := db.Take(&user, "id = ?", id).Error

	if user.ID == "" {
		return user, errors.New("user not found")
	} else {
		return user, err
	}
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.Save(&user).Error

	return user, err
}

func (u *UserRepositoryImpl) FindBy(db *gorm.DB, query string, value interface{}) (domain.User, error) {
	var user domain.User
	err := db.Take(&user, query, value).Error

	if user.ID == "" {
		return user, errors.New("user not found")
	} else {
		return user, err
	}
}

func (u *UserRepositoryImpl) ChangePassword(db *gorm.DB, id string, password string) error {
	err := db.Model(&domain.User{}).Where("id = ?", id).Update("password", password).Error

	return err
}
