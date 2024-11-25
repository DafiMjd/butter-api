package connectionmodel

import "butter/pkg/ctype"

type ConnectionEntity struct {
	FolloweeId ctype.NullString `gorm:"primaryKey"`
	FollowerId ctype.NullString `gorm:"primaryKey"`
}

func (c *ConnectionEntity) TableName() string {
	return "butter.connections"
}
