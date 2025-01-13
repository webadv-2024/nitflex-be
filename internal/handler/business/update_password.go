package business

import (
	"context"

	"nitflex/util"
)

func (b *business) UpdatePassword(ctx context.Context, username, password string) error {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}
	return b.repo.UpdatePassword(ctx, username, hashedPassword)
}
