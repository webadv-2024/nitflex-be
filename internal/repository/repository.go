package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateRefreshToken(ctx context.Context, params *UpdateRefreshTokenParams) error
}

type repository struct {
	*gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}
