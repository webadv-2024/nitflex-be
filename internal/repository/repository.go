package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Repository interface {
	GetUserByID(ctx context.Context, userID string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateRefreshToken(ctx context.Context, params *UpdateRefreshTokenParams) error
	UpdateUser(ctx context.Context, user *User) error

	// Movie methods
	GetMovieByID(ctx context.Context, movieID string) (*Movie, error)
	GetMoviesList(ctx context.Context, movieIDs []int) ([]*Movie, error)

	// Rating methods
	CreateRating(ctx context.Context, userID string, movieID string, rating int) (*Rating, error)
	GetRatingUser(ctx context.Context, userID string) ([]*Rating, error)
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
