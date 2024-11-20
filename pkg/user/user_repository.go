package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// Create implements UserRepository.
func (u *UserRepository) Create(user UserEntity) (UserEntity, error) {
	err := u.DB.Create(&user).Error

	return user, err
}

// Delete implements UserRepository.
func (u *UserRepository) Delete(user UserEntity) error {
	err := u.DB.Delete(&user).Error

	return err
}

// FindAll implements UserRepository.
func (u *UserRepository) FindAll() ([]UserEntity, error) {
	var users []UserEntity
	err := u.DB.Order("name asc").Find(&users).Error

	return users, err
}

// FindById implements UserRepository.
func (u *UserRepository) FindById(id string) (UserEntity, error) {
	var user UserEntity
	err := u.DB.Take(&user, "id = ?", id).Error

	if user.ID == "" {
		return user, errors.New("user not found")
	} else {
		return user, err
	}
}

// Update implements UserRepository.
func (u *UserRepository) Update(user UserEntity) (UserEntity, error) {
	err := u.DB.Save(&user).Error

	return user, err
}

func (u *UserRepository) FindBy(query string, value interface{}) (UserEntity, error) {
	var user UserEntity
	err := u.DB.Take(&user, query, value).Error

	if user.ID == "" {
		return user, errors.New("user not found")
	} else {
		return user, err
	}
}

func (u *UserRepository) ChangePassword(id string, password string) error {
	err := u.DB.Model(&UserEntity{}).Where("id = ?", id).Update("password", password).Error

	return err
}
