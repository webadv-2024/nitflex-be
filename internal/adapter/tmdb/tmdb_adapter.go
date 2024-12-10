package adapter

import (
	"context"
	"nitflex/internal/handler/models"
)

type TmdbAdapter interface {
	GetTrendingMovies(ctx context.Context, request *GetTrendingMoviesRequest) (*models.GetMoviesResponse, error)
	SearchMovies(ctx context.Context, request *SearchMoviesRequest) (*models.GetMoviesResponse, error)
}

type tmdbAdapter struct {
}

func NewTmdbAdapter() TmdbAdapter {
	return &tmdbAdapter{}
}
