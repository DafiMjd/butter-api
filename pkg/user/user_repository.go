package user

import (
	"butter/pkg/model/usermodel"
	"butter/pkg/pagination"
	"errors"
	"math"

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

func (u *UserRepository) FindAll(pgn *pagination.Pagination) ([]usermodel.UserEntity, error) {
	var users []usermodel.UserEntity
	if pgn.Sort == "" {
		pgn.Sort = "name asc"
	}

	err := u.DB.Scopes(pagination.Paginate(users, pgn, u.DB)).Find(&users).Error

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

func (u *UserRepository) FindAllByIds(ids []string, pgn *pagination.Pagination) ([]usermodel.UserEntity, error) {
	var users []usermodel.UserEntity
	err := u.DB.
		Scopes(pagination.PaginateOnly(
			pgn,
			u.DB,
		)).
		Where("id IN ?", ids).
		Find(&users).Error

	var totalDocs int64
	u.DB.Model(users).Where("id IN ?", ids).Count(&totalDocs)
	pgn.TotalDocs = totalDocs
	totalPages := int(math.Ceil(float64(totalDocs) / float64(pgn.GetLimit())))
	pgn.TotalPages = totalPages

	return users, err
}
