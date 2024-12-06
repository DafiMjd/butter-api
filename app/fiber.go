package app

import (
	"butter/middleware"
	"butter/pkg/connection"
	"butter/pkg/exception"
	"butter/pkg/post"
	"butter/pkg/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberHandlerSet struct {
	user.UserController
	post.PostController
	connection.ConnectionController
}

func NewFiber(hs FiberHandlerSet, am *middleware.AuthMiddleware) *fiber.App {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ErrorHandler,
		},
	)
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})
	v1 := app.Group("v1")
	butterGroup := v1.Group("butter")

	loginGroup := butterGroup.Group("login")
	loginGroup.Post("/username", hs.UserController.LoginWithUsername)
	loginGroup.Post("/email", hs.UserController.LoginWithEmail)

	butterGroup.Post("/signup", hs.UserController.Create)
	butterGroup.Get("/users", am.AuthenticateFiber(true), hs.UserController.FindAll)
	butterGroup.Get("/user/:userId", am.AuthenticateFiber(true), hs.UserController.FindById)
	butterGroup.Get("/user-username", am.AuthenticateFiber(true), hs.UserController.FindByUsername)
	butterGroup.Put("/user/:userId", am.AuthenticateFiber(false), hs.UserController.Update)
	butterGroup.Delete("/user/:userId", am.AuthenticateFiber(false), hs.UserController.Delete)
	butterGroup.Post("/refresh-token", am.AuthenticateRefreshToken(), hs.UserController.RefreshToken)
	butterGroup.Patch("/change-password/:userId", am.AuthenticateFiber(false), hs.UserController.ChangePassword)

	butterGroup.Post("/post", am.AuthenticateFiber(false), hs.PostController.Create)
	butterGroup.Get("/posts", am.AuthenticateFiber(true), hs.PostController.FindAll)
	butterGroup.Get("/post/:postId", am.AuthenticateFiber(true), hs.PostController.FindById)
	butterGroup.Patch("/post/:postId", am.AuthenticateFiber(false), hs.PostController.Update)
	butterGroup.Delete("/post/:postId", am.AuthenticateFiber(false), hs.PostController.Delete)

	butterGroup.Post("/follow", am.AuthenticateFiber(false), hs.ConnectionController.Follow)
	butterGroup.Delete("/unfollow", am.AuthenticateFiber(false), hs.ConnectionController.Unfollow)
	butterGroup.Get("/followers", am.AuthenticateFiber(true), hs.ConnectionController.FindAllFollowers)
	butterGroup.Get("/followings", am.AuthenticateFiber(true), hs.ConnectionController.FindAllFollowings)

	return app
}
