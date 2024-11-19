package app

import (
	"butter/middleware"
	"butter/pkg/exception"
	"butter/pkg/user/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberHandlerSet struct {
	UserController controller.UserController
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
	butterGroup.Put("/user/:userId", am.AuthenticateFiber(false), hs.UserController.Update)
	butterGroup.Delete("/user/:userId", am.AuthenticateFiber(false), hs.UserController.Delete)
	butterGroup.Post("/refresh-token", am.AuthenticateRefreshToken(), hs.UserController.RefreshToken)
	butterGroup.Patch("/change-password/:userId", am.AuthenticateFiber(false), hs.UserController.ChangePassword)

	return app
}
