package post

import (
	"butter/helper"
	"butter/pkg/exception"
	"butter/pkg/model/postmodel"
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

func (p *PostService) Create(request postmodel.PostCreateRequest) postmodel.PostResponse {
	user, err := p.UserRepository.FindById(request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	post := postmodel.PostEntity{
		ID:      uuid.New().String(),
		UserId:  request.UserId,
		Content: request.Content,
		User:    user,
	}
	createdPost, err := p.PostRepository.Create(post)
	helper.PanicIfError(err)

	return postmodel.ToPostResponse(createdPost)
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

func (p *PostService) FindAll(userId string) []postmodel.PostResponse {
	var posts []postmodel.PostEntity
	var err error
	if userId == "" {
		posts, err = p.PostRepository.FindAll()
	} else {
		posts, err = p.PostRepository.FindAllByUserId(userId)
	}
	helper.PanicIfError(err)

	return postmodel.ToPostResponses(posts)
}

func (p *PostService) FindById(id string) postmodel.PostResponse {
	post, err := p.PostRepository.FindById(id)
	helper.PanicIfError(err)

	return postmodel.ToPostResponse(post)
}

func (p *PostService) Update(
	request postmodel.PostUpdateRequest,
	loggedInUserId string,
) postmodel.PostResponse {
	post, err := p.PostRepository.FindById(request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	post.Content = request.Content

	if post.UserId != loggedInUserId {
		panic(exception.NewUnauthenticatedError("forbidden"))
	}

	var updatedPost postmodel.PostEntity

	updatedPost, err = p.PostRepository.Update(post)
	helper.PanicIfError(err)

	return postmodel.ToPostResponse(updatedPost)
}
