package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) FilterMovies(ctx context.Context, params *repository.FilterMoviesParams) ([]*repository.Movie, int, error) {
	if params.MaxRating == 0 {
		params.MaxRating = 10
	}

	return b.repo.FilterMovies(ctx, params)
}
