package connection

import (
	"butter/pkg/model/connectionmodel"
	"butter/pkg/model/usermodel"
	"fmt"

	"gorm.io/gorm"
)

type ConnectionRepository struct {
	DB *gorm.DB
}

func NewConnectionRepository(db *gorm.DB) *ConnectionRepository {
	return &ConnectionRepository{
		DB: db,
	}
}

func (c *ConnectionRepository) Follow(connection connectionmodel.ConnectionEntity) error {
	err := c.DB.Create(&connection).Error

	return err
}

func (c *ConnectionRepository) Unfollow(connection connectionmodel.ConnectionEntity) error {
	fmt.Println(connection)
	err := c.DB.Delete(&connection).Error

	return err
}

func (c *ConnectionRepository) FindAllFollowers(userId string) ([]usermodel.UserEntity, error) {
	var user usermodel.UserEntity
	err := c.DB.Preload("Followers").First(&user, "id = ?", userId).Error

	return user.Followers, err
}

func (c *ConnectionRepository) FindAllFollowings(userId string) ([]usermodel.UserEntity, error) {
	var user usermodel.UserEntity
	err := c.DB.Preload("Followings").First(&user, "id = ?", userId).Error

	return user.Followings, err
}
