package connectionmodel

type FollowRequest struct {
	FolloweeId string `json:"followeeId"`
	FollowerId string `json:"followerId"`
}
