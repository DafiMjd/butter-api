//go:build wireinject
// +build wireinject

package main

import (
	"butter/app"
	"butter/middleware"
	"butter/pkg/post"
	"butter/pkg/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func InitializedServer() *fiber.App {
	wire.Build(
		app.NewDb,
		middleware.NewAuthMiddleware,
		ProvideUser,
		ProvidePost,
		ProvideFiber,
	)
	return nil
}

var ProvideFiber = wire.NewSet(
	app.NewFiber,
	wire.Struct(new(app.FiberHandlerSet), "*"),
)

var ProvideUser = wire.NewSet(
	wire.Struct(new(user.UserRepository), "*"),
	wire.Struct(new(user.UserController), "*"),
	wire.Struct(new(user.UserService), "*"),
)

var ProvidePost = wire.NewSet(
	wire.Struct(new(post.PostRepository), "*"),
	wire.Struct(new(post.PostController), "*"),
	wire.Struct(new(post.PostService), "*"),
)
