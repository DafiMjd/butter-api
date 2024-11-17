//go:build wireinject
// +build wireinject

package main

import (
	"butter/app"
	"butter/feature/user/controller"
	"butter/feature/user/repository"
	"butter/feature/user/service"
	"butter/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func InitializedServer() *fiber.App {
	wire.Build(
		app.NewDb,
		repository.NewUserRepositoryImpl,
		service.NewUserServiceImpl,
		controller.NewUserControllerImpl,
		middleware.NewAuthMiddleware,
		ProvideFiber,
	)
	return nil
}

var ProvideFiber = wire.NewSet(
	app.NewFiber,
	wire.Struct(new(app.FiberHandlerSet), "*"),
)
