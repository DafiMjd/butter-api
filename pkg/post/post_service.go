package post

import (
	"butter/helper"
	"butter/pkg/exception"
	"butter/pkg/post/model/web"
	"butter/pkg/user"

	"github.com/google/uuid"
)

type PostService struct {
	PostRepository PostRepository
	UserRepository user.UserRepository
}

func NewPostService(
	postRepository PostRepository,
	userRepository user.UserRepository,
) *PostService {
	return &PostService{
		PostRepository: postRepository,
		UserRepository: userRepository,
	}
}

func (p *PostService) Create(request web.PostCreateRequest) web.PostResponse {
	user, err := p.UserRepository.FindById(request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	post := PostEntity{
		ID:      uuid.New().String(),
		UserId:  request.UserId,
		Content: request.Content,
		User:    user,
	}
	createdPost, err := p.PostRepository.Create(post)
	helper.PanicIfError(err)

	return ToPostResponse(createdPost)
}

func (p *PostService) Delete(id string, loggedInUserId string) {
	post, err := p.PostRepository.FindById(id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if post.UserId != loggedInUserId {
		panic(exception.NewUnauthenticatedError("forbidden"))
	}

	err = p.PostRepository.Delete(post)
	helper.PanicIfError(err)
}

func (p *PostService) FindAll(userId string) []web.PostResponse {
	var posts []PostEntity
	var err error
	if userId == "" {
		posts, err = p.PostRepository.FindAll()
	} else {
		posts, err = p.PostRepository.FindAllByUserId(userId)
	}
	helper.PanicIfError(err)

	return ToPostResponses(posts)
}

func (p *PostService) FindById(id string) web.PostResponse {
	post, err := p.PostRepository.FindById(id)
	helper.PanicIfError(err)

	return ToPostResponse(post)
}

func (p *PostService) Update(
	request web.PostUpdateRequest,
	loggedInUserId string,
) web.PostResponse {
	post, err := p.PostRepository.FindById(request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	post.Content = request.Content

	if post.UserId != loggedInUserId {
		panic(exception.NewUnauthenticatedError("forbidden"))
	}

	var updatedPost PostEntity

	updatedPost, err = p.PostRepository.Update(post)
	helper.PanicIfError(err)

	return ToPostResponse(updatedPost)
}
