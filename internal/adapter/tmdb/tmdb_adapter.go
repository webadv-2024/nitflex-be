package adapter

import (
	"context"

	"nitflex/internal/models"
)

type TmdbAdapter interface {
	GetTrendingMovies(ctx context.Context, timeWindow string) (*models.GetMoviesResponse, error)
}

type tmdbAdapter struct {
}

func NewTmdbAdapter() TmdbAdapter {
	return &tmdbAdapter{}
}
