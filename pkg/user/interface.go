package user

import (
	"butter/pkg/model/connectionmodel"
)

type IConnectionRepository interface {
	FindConnection(followerId string, followeeId string) (connectionmodel.ConnectionEntity, error)
	FindConnectionsIn(inQuery string) ([]connectionmodel.ConnectionEntity, error)
}
