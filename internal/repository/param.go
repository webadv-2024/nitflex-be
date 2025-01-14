package repository

import "time"

type UpdateRefreshTokenParams struct {
	UserId                string
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
}

type FilterMoviesParams struct {
	Genres         string
	MinRating      float64
	MaxRating      float64
	ReleaseDateGte string
	ReleaseDateLte string
	Actors         string
	Page           int
	PerPage        int
}
