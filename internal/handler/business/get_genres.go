package business

import (
	"context"

	"nitflex/internal/repository"
)

func (b *business) GetGenres(ctx context.Context) ([]*repository.Genre, error) {
	return b.repo.GetGenres(ctx)
}
