package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
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
	mongodb *mongo.Database
}

func NewRepository(db *gorm.DB, mongodb *mongo.Database) Repository {
	return &repository{
		DB:      db,
		mongodb: mongodb,
	}
}
