package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) SearchMovies(ctx context.Context, query string) ([]*repository.Movie, error) {
	return b.repo.SearchMoviesByQuery(ctx, query)
}
