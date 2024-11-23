package post

import (
	"butter/pkg/model/postmodel"
	"butter/pkg/pagination"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		DB: db,
	}
}

func (p *PostRepository) Create(post postmodel.PostEntity) (postmodel.PostEntity, error) {
	err := p.DB.Create(&post).Error

	return post, err
}

func (p *PostRepository) Delete(post postmodel.PostEntity) error {
	err := p.DB.Delete(&post).Error

	return err
}

func (p *PostRepository) FindAll(pgn *pagination.Pagination) ([]postmodel.PostEntity, error) {
	var posts []postmodel.PostEntity
	err := p.DB.
		Scopes(pagination.Paginate(posts, pgn, p.DB)).
		Preload("User").
		Order("created_at desc").
		Find(&posts).Error

	return posts, err
}

func (p *PostRepository) FindAllByUserId(userId string, pgn *pagination.Pagination) ([]postmodel.PostEntity, error) {

	var posts []postmodel.PostEntity
	err := p.DB.
		Scopes(pagination.Paginate(posts, pgn, p.DB)).
		Preload("User").
		Order("created_at desc").
		Find(&posts, "user_id = ?", userId).Error

	return posts, err
}

func (p *PostRepository) FindById(id string) (postmodel.PostEntity, error) {
	var post postmodel.PostEntity
	err := p.DB.Preload("User").Take(&post, "id = ?", id).Error

	return post, err
}

func (p *PostRepository) Update(post postmodel.PostEntity) (postmodel.PostEntity, error) {
	err := p.DB.
		Model(&post).
		Preload("User").
		Where("id = ?", post.ID).
		Update("content", post.Content).Error

	return post, err
}
