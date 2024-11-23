package connection

import (
	"butter/pkg/model/connectionmodel"

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
	err := c.DB.Delete(&connection).Error

	return err
}

func (c *ConnectionRepository) FindAllFollowerId(userId string) ([]string, error) {
	var connections []connectionmodel.ConnectionEntity
	err := c.DB.Find(&connections, "followee_id = ?", userId).Error
	ids := []string{}

	for _, connection := range connections {
		ids = append(ids, connection.FollowerId)
	}

	return ids, err
}

func (c *ConnectionRepository) FindAllFolloweeId(userId string) ([]string, error) {
	var connections []connectionmodel.ConnectionEntity
	err := c.DB.Find(&connections, "follower_id = ?", userId).Error
	ids := []string{}

	for _, connection := range connections {
		ids = append(ids, connection.FolloweeId)
	}

	return ids, err
}

func (c *ConnectionRepository) FindConnection(followerId string, followeeId string) (connectionmodel.ConnectionEntity, error) {
	var connection connectionmodel.ConnectionEntity
	err := c.DB.First(&connection, "follower_id = ? AND followee_id = ?", followerId, followeeId).Error

	return connection, err
}

func (c *ConnectionRepository) FindConnectionsIn(inQuery string) ([]connectionmodel.ConnectionEntity, error) {
	var res []connectionmodel.ConnectionEntity
	err := c.DB.Raw("SELECT * FROM butter.connections WHERE (followee_id, follower_id) IN " + inQuery).Scan(&res).Error

	return res, err
}
