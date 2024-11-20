//go:build wireinject
// +build wireinject

package main

import (
	"butter/app"
	"butter/middleware"
	"butter/pkg/connection"
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
		ProvideConnection,
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

var ProvideConnection = wire.NewSet(
	wire.Struct(new(connection.ConnectionRepository), "*"),
	wire.Struct(new(connection.ConnectionController), "*"),
	wire.Struct(new(connection.ConnectionService), "*"),
)
