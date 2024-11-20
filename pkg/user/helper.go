package user

import (
	"butter/pkg/user/model/web"
)

func ToUserResponse(user UserEntity) web.UserResponse {
	return web.UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		Birthdate: user.Birthdate,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []UserEntity) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
