package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) GetPopularMovies(ctx context.Context) ([]*repository.Movie, error) {
	return b.repo.GetPopularMovies(ctx)
}
