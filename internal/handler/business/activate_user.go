package business

import (
	"context"
)

func (b *business) ActivateUser(ctx context.Context, username string) error {
	err := b.repo.UpdateUserActivationStatus(ctx, username, true)
	if err != nil {
		return err
	}
	return nil
}
