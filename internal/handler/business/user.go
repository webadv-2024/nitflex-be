package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) GetUserByUsername(ctx context.Context, username string) (*repository.User, error) {
	user, err := b.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}