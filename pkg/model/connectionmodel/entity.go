package connectionmodel

type ConnectionEntity struct {
	FolloweeId string `gorm:"primaryKey"`
	FollowerId string `gorm:"primaryKey"`
}

func (c *ConnectionEntity) TableName() string {
	return "butter.connections"
}
