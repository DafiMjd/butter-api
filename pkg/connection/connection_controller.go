package connection

import (
	"butter/pkg/exception"
	"butter/pkg/model"
	"butter/pkg/model/connectionmodel"
	"butter/pkg/pagination"

	"github.com/gofiber/fiber/v2"
)

type ConnectionController struct {
	ConnectionService
}

func NewConnectionController(connectionService ConnectionService) *ConnectionController {
	return &ConnectionController{
		ConnectionService: connectionService,
	}
}

func (cn *ConnectionController) Follow(c *fiber.Ctx) error {
	request := connectionmodel.FollowRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	loggedInUserId, ok := c.Locals("user_id").(string)
	if !ok {
		panic(exception.NewUnauthenticatedError("unauthorized"))
	}

	request.FollowerId = loggedInUserId
	response := cn.ConnectionService.Follow(request)

	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: response,
		},
	}

	return c.JSON(webResponse)
}

func (cn *ConnectionController) Unfollow(c *fiber.Ctx) error {
	request := connectionmodel.FollowRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	loggedInUserId, ok := c.Locals("user_id").(string)
	if !ok {
		panic(exception.NewUnauthenticatedError("unauthorized"))
	}

	request.FollowerId = loggedInUserId
	response := cn.ConnectionService.Unfollow(request)

	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: response,
		},
	}

	return c.JSON(webResponse)
}

func (cn *ConnectionController) FindAllFollowers(c *fiber.Ctx) error {
	userId := c.Query("userId")

	pgn := pagination.Pagination{}
	err := c.QueryParser(&pgn)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	response := cn.ConnectionService.FindAllFollowers(userId, &pgn)
	pgn.Docs = response

	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data:   pgn,
	}

	return c.JSON(webResponse)
}

func (cn *ConnectionController) FindAllFollowings(c *fiber.Ctx) error {
	userId := c.Query("userId")
	pgn := pagination.Pagination{}
	err := c.QueryParser(&pgn)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	response := cn.ConnectionService.FindAllFollowings(userId, &pgn)
	pgn.Docs = response

	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data:   pgn,
	}

	return c.JSON(webResponse)
}
