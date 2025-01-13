package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) GetTrailers(ctx context.Context) ([]*repository.Trailer, error) {
	return b.repo.GetMovieTrailers(ctx)
}
