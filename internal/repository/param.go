package repository

import "time"

type UpdateRefreshTokenParams struct {
	UserId                int64
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
}
