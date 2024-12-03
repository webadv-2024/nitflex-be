package repository

import (
	"context"

	"gorm.io/gorm/clause"
)

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	var (
		user User
		err  = r.WithContext(ctx).First(&user, "username = ?", username).Error
	)

	return &user, err
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var (
		user User
		err  = r.WithContext(ctx).First(&user, "email = ?", email).Error
	)

	return &user, err
}

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	return r.WithContext(ctx).Table(User{}.TableName()).Create(user).Error
}

func (r *repository) UpdateRefreshToken(ctx context.Context, params *UpdateRefreshTokenParams) error {
	return r.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Table(User{}.TableName()).
		Where("id = ?", params.UserId).
		Updates(
			map[string]interface{}{
				"refresh_token":            params.RefreshToken,
				"refresh_token_expires_at": params.RefreshTokenExpiresAt,
			},
		).Error
}
