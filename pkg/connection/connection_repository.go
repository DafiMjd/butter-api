package connection

import (
	"butter/pkg/model/connectionmodel"
	"butter/pkg/pagination"
	"database/sql"

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
	err := c.DB.
		Where("follower_id = ?", connection.FollowerId).
		Where("followee_id = ?", connection.FolloweeId).
		Delete(&connection).Error

	return err
}

func (c *ConnectionRepository) FindAllFollowers(
	userId string,
	loggedInUserId string,
	pgn *pagination.Pagination,
) (*sql.Rows, error) {
	rows, err := c.DB.
		Table("butter.users u").
		Scopes(pagination.PaginateOnly(
			pgn,
			c.DB,
		)).
		Select("id", "username", "name", "email", "birthdate", "created_at", "updated_at", "b.followee_id", "b.follower_id").
		Joins("INNER JOIN butter.connections a ON u.id = a.follower_id").
		Joins("LEFT JOIN butter.connections b ON u.id = b.followee_id AND b.follower_id = ?", loggedInUserId).
		Where("a.followee_id = ?", userId).
		Rows()

	return rows, err
}

func (c *ConnectionRepository) FindAllFollowings(
	userId string,
	loggedInUserId string,
	pgn *pagination.Pagination,
) (*sql.Rows, error) {
	rows, err := c.DB.
		Table("butter.users u").
		Scopes(pagination.PaginateOnly(
			pgn,
			c.DB,
		)).
		Select("id", "username", "name", "email", "birthdate", "created_at", "updated_at", "b.follower_id").
		Joins("inner join butter.connections a on u.id = a.followee_id").
		Joins("LEFT JOIN butter.connections b ON u.id = b.followee_id AND b.follower_id = ?", loggedInUserId).
		Where("a.follower_id = ?", userId).
		Rows()

	return rows, err
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

func (c *ConnectionRepository) CountFollowers(followee_id string) (int64, error) {
	var totalDocs int64
	err := c.DB.Model(&connectionmodel.ConnectionEntity{}).Where("followee_id = ?", followee_id).Count(&totalDocs).Error

	return totalDocs, err
}

func (c *ConnectionRepository) CountFollowings(follower_id string) (int64, error) {
	var totalDocs int64
	err := c.DB.Model(&connectionmodel.ConnectionEntity{}).Where("follower_id = ?", follower_id).Count(&totalDocs).Error

	return totalDocs, err
}
