package repository

import "time"

type UpdateRefreshTokenParams struct {
	UserId                string
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
}
