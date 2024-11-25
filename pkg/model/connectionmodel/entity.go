package connectionmodel

import "butter/pkg/ctype"

type ConnectionEntity struct {
	FolloweeId       ctype.NullString `gorm:"primaryKey"`
	FolloweeUsername string
	FollowerId       ctype.NullString `gorm:"primaryKey"`
	FollowerUsername string
}

func (c *ConnectionEntity) TableName() string {
	return "butter.connections"
}
