package post

import (
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

func (p *PostRepository) Create(post PostEntity) (PostEntity, error) {
	err := p.DB.Create(&post).Error

	return post, err
}

func (p *PostRepository) Delete(post PostEntity) error {
	err := p.DB.Delete(&post).Error

	return err
}

func (p *PostRepository) FindAll() ([]PostEntity, error) {
	var posts []PostEntity
	err := p.DB.
		Preload("User").
		Order("created_at desc").
		Find(&posts).Error

	return posts, err
}

func (p *PostRepository) FindAllByUserId(userId string) ([]PostEntity, error) {

	var posts []PostEntity
	err := p.DB.
		Preload("User").
		Order("created_at desc").
		Find(&posts, "user_id = ?", userId).Error

	return posts, err
}

func (p *PostRepository) FindById(id string) (PostEntity, error) {
	var post PostEntity
	err := p.DB.Preload("User").Take(&post, "id = ?", id).Error

	return post, err
}

func (p *PostRepository) Update(post PostEntity) (PostEntity, error) {
	err := p.DB.
		Model(&post).
		Preload("User").
		Where("id = ?", post.ID).
		Update("content", post.Content).Error

	return post, err
}
