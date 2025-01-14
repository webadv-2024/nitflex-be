package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) SearchMovies(ctx context.Context, query string, page, perPage int) ([]*repository.Movie, int, error) {
	return b.repo.SearchMoviesByQuery(ctx, query, page, perPage)
}
