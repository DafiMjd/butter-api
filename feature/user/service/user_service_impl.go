package service

import (
	"butter/exception"
	"butter/feature/user/model/domain"
	"butter/feature/user/model/web"
	"butter/feature/user/repository"
	"butter/helper"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	// Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		// Validate:       validate,
	}
}

// Create implements UserService.
func (u *UserServiceImpl) Create(db *gorm.DB, request web.UserCreateRequest) web.LoginResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicIfError(err)

	newUser := domain.User{
		ID:        uuid.New().String(),
		Username:  request.Username,
		Password:  string(hash),
		Name:      request.Name,
		Email:     request.Email,
		Birthdate: request.Birthdate,
	}
	createdUser, err := u.UserRepository.Create(db, newUser)

	me, ok := err.(*mysql.MySQLError)
	if !ok {
		helper.PanicIfError(err)
	} else {
		if me.Number == 1062 {
			panic(exception.NewDuplicatedData("email or username already exists"))
		}
		helper.PanicIfError(err)
	}

	token := generateToken(createdUser, TokenExpiredTime)
	refreshToken := generateToken(createdUser, RefreshTokenExpiredTime)

	return web.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		UserResponse: helper.ToUserResponse(createdUser),
	}
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(db *gorm.DB, id string) {
	user, err := u.UserRepository.FindById(db, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = u.UserRepository.Delete(db, user)
	helper.PanicIfError(err)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll(db *gorm.DB) []web.UserResponse {
	users, err := u.UserRepository.FindAll(db)
	helper.PanicIfError(err)

	return helper.ToUserResponses(users)
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(db *gorm.DB, id string) web.UserResponse {
	user, err := u.UserRepository.FindById(db, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (u *UserServiceImpl) Update(db *gorm.DB, request web.UserUpdateRequest) web.UserResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	user, err := u.UserRepository.FindById(db, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Username = request.Username
	user.Name = request.Name
	user.Birthdate = request.Birthdate

	var updatedUser = domain.User{}

	updatedUser, err = u.UserRepository.Update(db, user)
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		helper.PanicIfError(err)
	} else {
		if me.Number == 1062 {
			panic(exception.NewDuplicatedData("username already exists"))
		}
		helper.PanicIfError(err)
	}

	return helper.ToUserResponse(updatedUser)
}

func (u *UserServiceImpl) LoginWithUsername(db *gorm.DB, request web.LoginRequest) web.LoginResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	if request.Username == "" {
		panic(exception.NewBadRequestError("Username is required"))
	}

	return login(u.UserRepository, db, "username = ?", request.Username, request.Password)
}

func (u *UserServiceImpl) LoginWithEmail(db *gorm.DB, request web.LoginRequest) web.LoginResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	if request.Email == "" {
		panic(exception.NewBadRequestError("Email is required"))
	}

	return login(u.UserRepository, db, "email = ?", request.Email, request.Password)
}

func login(repo repository.UserRepository, db *gorm.DB, query string, value string, password string) web.LoginResponse {
	if password == "" {
		panic(exception.NewBadRequestError("Password is required"))
	}

	user, err := repo.FindBy(db, query, value)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		panic(exception.NewBadRequestError("Password is incorrect"))
	}

	token := generateToken(user, TokenExpiredTime)
	refreshToken := generateToken(user, RefreshTokenExpiredTime)

	return web.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		UserResponse: helper.ToUserResponse(user),
	}
}

func generateToken(user domain.User, exp time.Duration) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(exp).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	helper.PanicIfError(err)

	return tokenString
}

// RefreshToken implements UserService.
func (u *UserServiceImpl) RefreshToken(user domain.User) web.LoginResponse {
	token := generateToken(user, TokenExpiredTime)
	refreshToken := generateToken(user, RefreshTokenExpiredTime)

	return web.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		UserResponse: helper.ToUserResponse(user),
	}
}

func (u *UserServiceImpl) ChangePassword(db *gorm.DB, request web.ChangePasswordRequest) {
	user, err := u.UserRepository.FindById(db, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.OldPassword == "" {
		panic(exception.NewBadRequestError("Old Password is required"))
	}
	if request.NewPassword == "" {
		panic(exception.NewBadRequestError("New Password is required"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		panic(exception.NewBadRequestError("Old Password is incorrect"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.NewPassword))
	if err == nil {
		panic(exception.NewBadRequestError("New password must not be the same as old password"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 10)
	helper.PanicIfError(err)

	err = u.UserRepository.ChangePassword(db, request.ID, string(hash))
	helper.PanicIfError(err)
}
