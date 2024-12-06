package connectionmodel

import (
	"github.com/google/uuid"
)

type ConnectionEntity struct {
	FolloweeId       uuid.UUID `gorm:"primaryKey"`
	FolloweeUsername string
	FollowerId       uuid.UUID `gorm:"primaryKey"`
	FollowerUsername string
}

func (c *ConnectionEntity) TableName() string {
	return "butter.connections"
}
