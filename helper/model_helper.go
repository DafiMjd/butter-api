package helper

import (
	"butter/pkg/user/model/domain"
	"butter/pkg/user/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
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

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
