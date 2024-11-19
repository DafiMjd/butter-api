package controller

import (
	"butter/model"
	"butter/pkg/exception"
	"butter/pkg/user/model/domain"
	"butter/pkg/user/model/web"
	"butter/pkg/user/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserControllerImpl struct {
	UserService service.UserService
	DB          *gorm.DB
}

func NewUserControllerImpl(userService service.UserService, db *gorm.DB) UserController {
	return &UserControllerImpl{
		UserService: userService,
		DB:          db,
	}
}

// Create implements UserController.
func (u *UserControllerImpl) Create(c *fiber.Ctx) error {
	userCreateRequest := web.UserCreateRequest{}
	err := c.BodyParser(&userCreateRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	userResponse := u.UserService.Create(u.DB, userCreateRequest)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: userResponse,
		},
	}

	return c.JSON(webResponse)
}

// Delete implements UserController.
func (u *UserControllerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("userId")
	checkUserId(c, id)

	u.UserService.Delete(u.DB, id)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
	}

	return c.JSON(webResponse)
}

// FindAll implements UserController.
func (u *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	users := u.UserService.FindAll(u.DB)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.MultiDocs{
			Docs: users,
		},
	}

	return c.JSON(webResponse)
}

// FindById implements UserController.
func (u *UserControllerImpl) FindById(c *fiber.Ctx) error {
	id := c.Params("userId")
	user := u.UserService.FindById(u.DB, id)

	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: user,
		},
	}

	return c.JSON(webResponse)
}

// LoginWithEmail implements UserController.
func (u *UserControllerImpl) LoginWithEmail(c *fiber.Ctx) error {
	loginRequest := web.LoginRequest{}
	err := c.BodyParser(&loginRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	user := u.UserService.LoginWithEmail(u.DB, loginRequest)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: user,
		},
	}

	return c.JSON(webResponse)
}

// LoginWithUsername implements UserController.
func (u *UserControllerImpl) LoginWithUsername(c *fiber.Ctx) error {
	loginRequest := web.LoginRequest{}
	err := c.BodyParser(&loginRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	user := u.UserService.LoginWithUsername(u.DB, loginRequest)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: user,
		},
	}

	return c.JSON(webResponse)
}

// Update implements UserController.
func (u *UserControllerImpl) Update(c *fiber.Ctx) error {
	id := c.Params("userId")
	checkUserId(c, id)

	request := web.UserUpdateRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	request.ID = id

	user := u.UserService.Update(u.DB, request)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: user,
		},
	}

	return c.JSON(webResponse)
}

func checkUserId(c *fiber.Ctx, paramUserId string) {
	loggedInUserId, ok := c.Locals("user_id").(string)
	if !ok || loggedInUserId != paramUserId {
		panic(exception.NewUnauthenticatedError("unauthorized"))
	}
}

func (u *UserControllerImpl) RefreshToken(c *fiber.Ctx) error {
	id, ok := c.Locals("user_id").(string)
	if !ok {
		panic(exception.NewUnauthenticatedError("unauthorized"))
	}

	response := u.UserService.RefreshToken(domain.User{ID: id})
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: response,
		},
	}

	return c.JSON(webResponse)
}

func (u *UserControllerImpl) ChangePassword(c *fiber.Ctx) error {
	id := c.Params("userId")
	checkUserId(c, id)

	request := web.ChangePasswordRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	request.ID = id

	u.UserService.ChangePassword(u.DB, request)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
	}

	return c.JSON(webResponse)
}
