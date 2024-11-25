package connection

import (
	"butter/helper"
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
	err := c.DB.
		Where("follower_id = ?", connection.FollowerId).
		Where("followee_id = ?", connection.FolloweeId).
		Delete(&connection).Error

	return err
}

func (c *ConnectionRepository) FindAllFollowers(userId string) ([]usermodel.UserEntity, error) {
	rows, err := c.DB.
		Table("users").
		Select("id", "username", "name", "email", "birthdate", "created_at", "updated_at", "b.followee_id", "b.follower_id").
		Joins("INNER JOIN connections a ON users.id = a.follower_id").
		Joins("LEFT JOIN connections b ON users.id = b.followee_id AND b.follower_id = ?", userId).
		Where("a.followee_id = ?", userId).
		Rows()
	helper.PanicIfError(err)
	defer rows.Close()

	var users []usermodel.UserEntity
	for rows.Next() {
		user := usermodel.UserEntity{}
		conn := connectionmodel.ConnectionEntity{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Name,
			&user.Birthdate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&conn.FolloweeId,
			&conn.FollowerId,
		)
		user.IsFollowed = conn.FolloweeId.Valid
		helper.PanicIfError(err)
		fmt.Println("new")
		fmt.Println(user)
		users = append(users, user)
	}

	return users, err
}

func (c *ConnectionRepository) FindAllFollowings(userId string) ([]usermodel.UserEntity, error) {
	rows, err := c.DB.
		Table("users").
		Select("id", "username", "name", "email", "birthdate", "created_at", "updated_at").
		Joins("inner join connections on users.id = connections.followee_id").
		Where("connections.follower_id = ?", userId).
		Rows()
	helper.PanicIfError(err)
	defer rows.Close()

	var users []usermodel.UserEntity
	for rows.Next() {
		user := usermodel.UserEntity{
			IsFollowed: true,
		}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Birthdate, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users, err
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
