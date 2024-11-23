package connection

import (
	"butter/helper"
	"butter/pkg/exception"
	"butter/pkg/model/connectionmodel"
	"butter/pkg/model/usermodel"
	"butter/pkg/pagination"
	"butter/pkg/user"
)

type ConnectionService struct {
	ConnectionRepository ConnectionRepository
	UserRepository       user.UserRepository
}

func NewConnectionService(
	connectionRepository ConnectionRepository,
	userRepository user.UserRepository,
) *ConnectionService {
	return &ConnectionService{
		ConnectionRepository: connectionRepository,
		UserRepository:       userRepository,
	}
}

func (c *ConnectionService) Follow(request connectionmodel.FollowRequest) usermodel.UserResponse {
	users, err := c.UserRepository.FindAllByIds(
		[]string{request.FollowerId, request.FolloweeId},
		&pagination.Pagination{
			Limit: 2,
		},
	)

	if err != nil || len(users) != 2 {
		panic(exception.NewNotFoundError(err.Error()))
	}

	connection := connectionmodel.ConnectionEntity(request)

	var followee usermodel.UserEntity
	for _, user := range users {
		if user.ID == request.FolloweeId {
			followee = user
		}
	}
	followee.IsFollowed = true

	err = c.ConnectionRepository.Follow(connection)
	helper.PanicIfError(err)

	return usermodel.ToUserResponse(followee)
}

func (c *ConnectionService) Unfollow(request connectionmodel.FollowRequest) usermodel.UserResponse {
	users, err := c.UserRepository.
		FindAllByIds(
			[]string{request.FollowerId, request.FolloweeId},
			&pagination.Pagination{
				Limit: 2,
			},
		)

	if err != nil || len(users) != 2 {
		panic(exception.NewNotFoundError(err.Error()))
	}

	connection := connectionmodel.ConnectionEntity(request)

	var followee usermodel.UserEntity
	for _, user := range users {
		if user.ID == request.FolloweeId {
			followee = user
		}
	}

	err = c.ConnectionRepository.Unfollow(connection)
	helper.PanicIfError(err)

	return usermodel.ToUserResponse(followee)
}

func (c *ConnectionService) FindAllFollowers(userId string, pgn *pagination.Pagination) []usermodel.UserResponse {
	followersId, err := c.ConnectionRepository.FindAllFollowerId(userId)
	helper.PanicIfError(err)

	users, err := c.UserRepository.FindAllByIds(followersId, pgn)
	helper.PanicIfError(err)

	return usermodel.ToUserResponses(users)
}

func (c *ConnectionService) FindAllFollowings(userId string, pgn *pagination.Pagination) []usermodel.UserResponse {
	followeesId, err := c.ConnectionRepository.FindAllFolloweeId(userId)
	helper.PanicIfError(err)

	users, err := c.UserRepository.FindAllByIds(followeesId, pgn)
	helper.PanicIfError(err)

	return usermodel.ToUserResponses(users)
}
