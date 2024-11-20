// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from injector.go:

func InitializedServer() *fiber.App {
	db := app.NewDb()
	userRepository := user.UserRepository{
		DB: db,
	}
	userService := user.UserService{
		UserRepository: userRepository,
	}
	userController := user.UserController{
		UserService: userService,
	}
	postRepository := post.PostRepository{
		DB: db,
	}
	postService := post.PostService{
		PostRepository: postRepository,
		UserRepository: userRepository,
	}
	postController := post.PostController{
		PostService: postService,
	}
	connectionRepository := connection.ConnectionRepository{
		DB: db,
	}
	connectionService := connection.ConnectionService{
		ConnectionRepository: connectionRepository,
		UserRepository:       userRepository,
	}
	connectionController := connection.ConnectionController{
		ConnectionService: connectionService,
	}
	fiberHandlerSet := app.FiberHandlerSet{
		UserController:       userController,
		PostController:       postController,
		ConnectionController: connectionController,
	}
	authMiddleware := middleware.NewAuthMiddleware(userService, db)
	fiberApp := app.NewFiber(fiberHandlerSet, authMiddleware)
	return fiberApp
}

// injector.go:

var ProvideFiber = wire.NewSet(app.NewFiber, wire.Struct(new(app.FiberHandlerSet), "*"))

var ProvideUser = wire.NewSet(wire.Struct(new(user.UserRepository), "*"), wire.Struct(new(user.UserController), "*"), wire.Struct(new(user.UserService), "*"))

var ProvidePost = wire.NewSet(wire.Struct(new(post.PostRepository), "*"), wire.Struct(new(post.PostController), "*"), wire.Struct(new(post.PostService), "*"))

var ProvideConnection = wire.NewSet(wire.Struct(new(connection.ConnectionRepository), "*"), wire.Struct(new(connection.ConnectionController), "*"), wire.Struct(new(connection.ConnectionService), "*"))
