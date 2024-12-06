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

	connection := connectionmodel.ConnectionEntity{
		FolloweeId: helper.StringToUUID(request.FolloweeId),
		FollowerId: helper.StringToUUID(request.FollowerId),
	}

	var followee usermodel.UserEntity
	for _, user := range users {
		if user.ID.String() == request.FolloweeId {
			followee = user
			connection.FolloweeUsername = user.Username
		}

		if user.ID.String() == request.FollowerId {
			connection.FollowerUsername = user.Username
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

	var followee usermodel.UserEntity
	for _, user := range users {
		if user.ID.String() == request.FolloweeId {
			followee = user
		}
	}

	connection := connectionmodel.ConnectionEntity{
		FolloweeId: helper.StringToUUID(request.FolloweeId),
		FollowerId: helper.StringToUUID(request.FollowerId),
	}

	err = c.ConnectionRepository.Unfollow(connection)
	helper.PanicIfError(err)

	return usermodel.ToUserResponse(followee)
}

func (c *ConnectionService) FindAllFollowers(userId string, pgn *pagination.Pagination) []usermodel.UserResponse {
	rows, err := c.ConnectionRepository.FindAllFollowers(userId, pgn)
	helper.PanicIfError(err)

	defer rows.Close()

	var users []usermodel.UserEntity
	for rows.Next() {
		user := usermodel.UserEntity{}
		conn := connectionmodel.ConnectionEntity{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Name,
			&user.Birthdate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&conn.FolloweeId,
			&conn.FollowerId,
		)

		user.IsFollowed = helper.IsUUIDValid(conn.FollowerId.String())
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return usermodel.ToUserResponses(users)
}

func (c *ConnectionService) FindAllFollowings(userId string, pgn *pagination.Pagination) []usermodel.UserResponse {
	rows, err := c.ConnectionRepository.FindAllFollowings(userId, pgn)
	helper.PanicIfError(err)

	defer rows.Close()

	var users []usermodel.UserEntity
	for rows.Next() {
		user := usermodel.UserEntity{
			IsFollowed: true,
		}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Name,
			&user.Birthdate,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return usermodel.ToUserResponses(users)
}
