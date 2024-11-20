package user

import (
	"butter/helper"
	"butter/pkg/exception"
	"butter/pkg/user/model/web"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository UserRepository
	// Validate       *validator.Validate
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
		// Validate:       validate,
	}
}

// Create implements UserService.
func (u *UserService) Create(request web.UserCreateRequest) web.LoginResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicIfError(err)

	newUser := UserEntity{
		ID:        uuid.New().String(),
		Username:  request.Username,
		Password:  string(hash),
		Name:      request.Name,
		Email:     request.Email,
		Birthdate: request.Birthdate,
	}
	createdUser, err := u.UserRepository.Create(newUser)

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
		UserResponse: ToUserResponse(createdUser),
	}
}

// Delete implements UserService.
func (u *UserService) Delete(id string) {
	user, err := u.UserRepository.FindById(id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = u.UserRepository.Delete(user)
	helper.PanicIfError(err)
}

// FindAll implements UserService.
func (u *UserService) FindAll() []web.UserResponse {
	users, err := u.UserRepository.FindAll()
	helper.PanicIfError(err)

	return ToUserResponses(users)
}

// FindById implements UserService.
func (u *UserService) FindById(id string) web.UserResponse {
	user, err := u.UserRepository.FindById(id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToUserResponse(user)
}

func (u *UserService) Update(request web.UserUpdateRequest) web.UserResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	user, err := u.UserRepository.FindById(request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Username = request.Username
	user.Name = request.Name
	user.Birthdate = request.Birthdate

	var updatedUser = UserEntity{}

	updatedUser, err = u.UserRepository.Update(user)
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		helper.PanicIfError(err)
	} else {
		if me.Number == 1062 {
			panic(exception.NewDuplicatedData("username already exists"))
		}
		helper.PanicIfError(err)
	}

	return ToUserResponse(updatedUser)
}

func (u *UserService) LoginWithUsername(request web.LoginRequest) web.LoginResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	if request.Username == "" {
		panic(exception.NewBadRequestError("Username is required"))
	}

	return login(u.UserRepository, "username = ?", request.Username, request.Password)
}

func (u *UserService) LoginWithEmail(request web.LoginRequest) web.LoginResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	if request.Email == "" {
		panic(exception.NewBadRequestError("Email is required"))
	}

	return login(u.UserRepository, "email = ?", request.Email, request.Password)
}

func login(repo UserRepository, query string, value string, password string) web.LoginResponse {
	if password == "" {
		panic(exception.NewBadRequestError("Password is required"))
	}

	user, err := repo.FindBy(query, value)
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
		UserResponse: ToUserResponse(user),
	}
}

func generateToken(user UserEntity, exp time.Duration) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(exp).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	helper.PanicIfError(err)

	return tokenString
}

// RefreshToken implements UserService.
func (u *UserService) RefreshToken(user UserEntity) web.LoginResponse {
	token := generateToken(user, TokenExpiredTime)
	refreshToken := generateToken(user, RefreshTokenExpiredTime)

	return web.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		UserResponse: ToUserResponse(user),
	}
}

func (u *UserService) ChangePassword(request web.ChangePasswordRequest) {
	user, err := u.UserRepository.FindById(request.ID)
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

	err = u.UserRepository.ChangePassword(request.ID, string(hash))
	helper.PanicIfError(err)
}
