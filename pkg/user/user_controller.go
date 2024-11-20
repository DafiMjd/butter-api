package user

import (
	"butter/pkg/exception"
	"butter/pkg/model"
	"butter/pkg/model/usermodel"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Create implements UserController.
func (u *UserController) Create(c *fiber.Ctx) error {
	userCreateRequest := usermodel.UserCreateRequest{}
	err := c.BodyParser(&userCreateRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	userResponse := u.UserService.Create(userCreateRequest)
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
func (u *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("userId")
	checkUserId(c, id)

	u.UserService.Delete(id)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
	}

	return c.JSON(webResponse)
}

// FindAll implements UserController.
func (u *UserController) FindAll(c *fiber.Ctx) error {
	users := u.UserService.FindAll()
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
func (u *UserController) FindById(c *fiber.Ctx) error {
	id := c.Params("userId")
	user := u.UserService.FindById(id)

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
func (u *UserController) LoginWithEmail(c *fiber.Ctx) error {
	loginRequest := usermodel.LoginRequest{}
	err := c.BodyParser(&loginRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	user := u.UserService.LoginWithEmail(loginRequest)
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
func (u *UserController) LoginWithUsername(c *fiber.Ctx) error {
	loginRequest := usermodel.LoginRequest{}
	err := c.BodyParser(&loginRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	user := u.UserService.LoginWithUsername(loginRequest)
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
func (u *UserController) Update(c *fiber.Ctx) error {
	id := c.Params("userId")
	checkUserId(c, id)

	request := usermodel.UserUpdateRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	request.ID = id

	user := u.UserService.Update(request)
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

func (u *UserController) RefreshToken(c *fiber.Ctx) error {
	id, ok := c.Locals("user_id").(string)
	if !ok {
		panic(exception.NewUnauthenticatedError("unauthorized"))
	}

	response := u.UserService.RefreshToken(usermodel.UserEntity{ID: id})
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: response,
		},
	}

	return c.JSON(webResponse)
}

func (u *UserController) ChangePassword(c *fiber.Ctx) error {
	id := c.Params("userId")
	checkUserId(c, id)

	request := usermodel.ChangePasswordRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	request.ID = id

	u.UserService.ChangePassword(request)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
	}

	return c.JSON(webResponse)
}
