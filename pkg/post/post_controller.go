package post

import (
	"butter/pkg/exception"
	"butter/pkg/model"
	"butter/pkg/model/postmodel"
	"butter/pkg/pagination"

	"github.com/gofiber/fiber/v2"
)

type PostController struct {
	PostService PostService
}

func NewPostController(postService PostService) *PostController {
	return &PostController{
		PostService: postService,
	}
}

func (p *PostController) Create(c *fiber.Ctx) error {
	postCreateRequest := postmodel.PostCreateRequest{}
	err := c.BodyParser(&postCreateRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	loggedInUserId, ok := c.Locals("user_id").(string)
	if !ok {
		panic(exception.NewUnauthenticatedError("unauthorized"))
	}
	postCreateRequest.UserId = loggedInUserId

	userResponse := p.PostService.Create(postCreateRequest)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: userResponse,
		},
	}

	return c.JSON(webResponse)
}

func (p *PostController) Delete(c *fiber.Ctx) error {
	id := c.Params("postId")

	loggedInUserId, ok := c.Locals("user_id").(string)
	if !ok {
		panic(exception.NewUnauthenticatedError("unauthorized"))
	}

	p.PostService.Delete(id, loggedInUserId)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
	}

	return c.JSON(webResponse)
}

func (p *PostController) FindAll(c *fiber.Ctx) error {
	userId := c.Query("userId")

	pgn := pagination.Pagination{}
	err := c.QueryParser(&pgn)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	posts := p.PostService.FindAll(userId, &pgn)
	pgn.Docs = posts

	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data:   pgn,
	}

	return c.JSON(webResponse)
}

func (p *PostController) FindById(c *fiber.Ctx) error {
	id := c.Params("postId")

	post := p.PostService.FindById(id)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: post,
		},
	}

	return c.JSON(webResponse)
}

func (p *PostController) Update(c *fiber.Ctx) error {
	id := c.Params("postId")

	request := postmodel.PostUpdateRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	request.ID = id
	loggedInUserId, ok := c.Locals("user_id").(string)
	if !ok {
		panic(exception.NewUnauthenticatedError("unauthorized"))
	}

	user := p.PostService.Update(request, loggedInUserId)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "success",
		Data: model.SingleDoc{
			Doc: user,
		},
	}

	return c.JSON(webResponse)
}
