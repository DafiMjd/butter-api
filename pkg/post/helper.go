package post

import (
	"butter/pkg/post/model/web"
	"butter/pkg/user"
)

func ToPostResponse(post PostEntity) web.PostResponse {
	return web.PostResponse{
		ID:           post.ID,
		UserId:       post.UserId,
		Content:      post.Content,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
		UserResponse: user.ToUserResponse(post.User),
	}
}

func ToPostResponses(posts []PostEntity) []web.PostResponse {
	var userResponses []web.PostResponse
	for _, post := range posts {
		userResponses = append(userResponses, ToPostResponse(post))
	}

	return userResponses
}
