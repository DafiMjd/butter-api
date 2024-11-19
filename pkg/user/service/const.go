package service

import "time"

var (
	TokenExpiredTime        = time.Hour * 24 * 30
	RefreshTokenExpiredTime = time.Hour * 24 * 30 * 3
)
