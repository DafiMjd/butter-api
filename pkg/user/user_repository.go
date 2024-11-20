package user

import (
	"butter/pkg/model/usermodel"
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

func (u *UserRepository) Create(user usermodel.UserEntity) (usermodel.UserEntity, error) {
	err := u.DB.Create(&user).Error

	return user, err
}

func (u *UserRepository) Delete(user usermodel.UserEntity) error {
	err := u.DB.Delete(&user).Error

	return err
}

func (u *UserRepository) FindAll() ([]usermodel.UserEntity, error) {
	var users []usermodel.UserEntity
	err := u.DB.Order("name asc").Find(&users).Error

	return users, err
}

func (u *UserRepository) FindById(id string) (usermodel.UserEntity, error) {
	var user usermodel.UserEntity
	err := u.DB.Take(&user, "id = ?", id).Error

	if user.ID == "" {
		return user, errors.New("user not found")
	} else {
		return user, err
	}
}

func (u *UserRepository) Update(user usermodel.UserEntity) (usermodel.UserEntity, error) {
	err := u.DB.Save(&user).Error

	return user, err
}

func (u *UserRepository) FindBy(query string, value interface{}) (usermodel.UserEntity, error) {
	var user usermodel.UserEntity
	err := u.DB.Take(&user, query, value).Error

	if user.ID == "" {
		return user, errors.New("user not found")
	} else {
		return user, err
	}
}

func (u *UserRepository) ChangePassword(id string, password string) error {
	err := u.DB.Model(&usermodel.UserEntity{}).Where("id = ?", id).Update("password", password).Error

	return err
}

func (u *UserRepository) FindAllByIds(ids []string) ([]usermodel.UserEntity, error) {
	var users []usermodel.UserEntity
	err := u.DB.Where("id IN ?", ids).Order("name asc").Find(&users).Error

	return users, err
}
