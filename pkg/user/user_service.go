package user

import (
	"butter/helper"
	"butter/pkg/exception"
	"butter/pkg/model/connectionmodel"
	"butter/pkg/model/usermodel"
	"butter/pkg/pagination"
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
	ConnectionRepository IConnectionRepository
}

func NewUserService(userRepository UserRepository, connectionRepository IConnectionRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
		// Validate:       validate,
		ConnectionRepository: connectionRepository,
	}
}

// Create implements UserService.
func (u *UserService) Create(request usermodel.UserCreateRequest) usermodel.LoginResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicIfError(err)

	newUser := usermodel.UserEntity{
		ID:        uuid.New(),
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

	return usermodel.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		UserResponse: usermodel.ToUserResponse(createdUser),
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
func (u *UserService) FindAll(loggedInId string, pgn *pagination.Pagination) []usermodel.UserResponse {
	users, err := u.UserRepository.FindAll(pgn)
	helper.PanicIfError(err)

	var connections []connectionmodel.ConnectionEntity
	if loggedInId != "" && len(users) > 0 {
		inQuery := ""
		for index, user := range users {
			inQuery += "(\"" + user.ID.String() + "\", \"" + loggedInId + "\")"

			if index < len(users)-1 {
				inQuery += ", "
			}
		}
		// inQuery += ")"

		connections, _ = u.ConnectionRepository.FindConnectionsIn(inQuery)

		conMap := make(map[string]string)
		for _, con := range connections {
			conMap[con.FolloweeId.String()] = con.FollowerId.String()
		}

		for index, user := range users {
			if _, ok := conMap[user.ID.String()]; ok {
				user.IsFollowed = true
				users[index] = user
			}
		}
	}

	return usermodel.ToUserResponses(users)
}

// FindById implements UserService.
func (u *UserService) FindById(id string, loggedInId string) usermodel.UserResponse {
	user, err := u.UserRepository.FindById(id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if loggedInId != "" && loggedInId != id {
		_, err := u.ConnectionRepository.FindConnection(loggedInId, id)

		user.IsFollowed = err == nil
	}

	return usermodel.ToUserResponse(user)
}

func (u *UserService) Update(request usermodel.UserUpdateRequest) usermodel.UserResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	user, err := u.UserRepository.FindById(request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Username = request.Username
	user.Name = request.Name
	user.Birthdate = request.Birthdate

	var updatedUser = usermodel.UserEntity{}

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

	return usermodel.ToUserResponse(updatedUser)
}

func (u *UserService) LoginWithUsername(request usermodel.LoginRequest) usermodel.LoginResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	if request.Username == "" {
		panic(exception.NewBadRequestError("Username is required"))
	}

	return login(u.UserRepository, "username = ?", request.Username, request.Password)
}

func (u *UserService) LoginWithEmail(request usermodel.LoginRequest) usermodel.LoginResponse {
	// err := u.Validate.Struct(request)
	// helper.PanicIfError(err)

	if request.Email == "" {
		panic(exception.NewBadRequestError("Email is required"))
	}

	return login(u.UserRepository, "email = ?", request.Email, request.Password)
}

func login(repo UserRepository, query string, value string, password string) usermodel.LoginResponse {
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

	return usermodel.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		UserResponse: usermodel.ToUserResponse(user),
	}
}

func generateToken(user usermodel.UserEntity, exp time.Duration) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(exp).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	helper.PanicIfError(err)

	return tokenString
}

// RefreshToken implements UserService.
func (u *UserService) RefreshToken(user usermodel.UserEntity) usermodel.LoginResponse {
	token := generateToken(user, TokenExpiredTime)
	refreshToken := generateToken(user, RefreshTokenExpiredTime)

	return usermodel.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		UserResponse: usermodel.ToUserResponse(user),
	}
}

func (u *UserService) ChangePassword(request usermodel.ChangePasswordRequest) {
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
