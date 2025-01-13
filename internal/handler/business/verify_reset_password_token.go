package business

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

	"nitflex/internal/repository"
	"nitflex/util"
)

func (b *business) VerifyResetPasswordToken(ctx context.Context, token string) (*repository.User, error) {
	user, err := b.repo.GetUserByResetPasswordToken(ctx, token)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, util.NewError("Invalid or expired token")
		}
		return nil, err
	}

	// Kiểm tra xem token có còn hiệu lực không
	if user.ResetPasswordToken != token {
		return nil, util.NewError("Invalid token")
	}

	return user, nil
}
