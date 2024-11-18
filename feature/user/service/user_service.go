package service

import (
	"butter/feature/user/model/domain"
	"butter/feature/user/model/web"

	"gorm.io/gorm"
)

type UserService interface {
	Create(db *gorm.DB, request web.UserCreateRequest) web.LoginResponse
	Update(db *gorm.DB, request web.UserUpdateRequest) web.UserResponse
	Delete(db *gorm.DB, id string)
	FindById(db *gorm.DB, id string) web.UserResponse
	FindAll(db *gorm.DB) []web.UserResponse
	LoginWithUsername(db *gorm.DB, request web.LoginRequest) web.LoginResponse
	LoginWithEmail(db *gorm.DB, request web.LoginRequest) web.LoginResponse
	RefreshToken(user domain.User) web.LoginResponse
}
