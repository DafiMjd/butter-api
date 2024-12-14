package user

import (
	"butter/pkg/model/connectionmodel"
)

type IConnectionRepository interface {
	FindConnection(followerId string, followeeId string) (connectionmodel.ConnectionEntity, error)
	FindConnectionsIn(inQuery string) ([]connectionmodel.ConnectionEntity, error)
	CountFollowers(followee_id string) (int64, error)
	CountFollowings(follower_id string) (int64, error)
}
