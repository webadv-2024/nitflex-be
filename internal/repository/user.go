package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.mongodb.Collection(User{}.TableName()).
		FindOne(ctx, bson.M{"username": username}).
		Decode(&user)
	return &user, err
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.mongodb.Collection(User{}.TableName()).
		FindOne(ctx, bson.M{"email": email}).
		Decode(&user)
	return &user, err
}

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	collection := r.mongodb.Collection(User{}.TableName())
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (r *repository) UpdateRefreshToken(ctx context.Context, params *UpdateRefreshTokenParams) error {
	filter := bson.M{"_id": params.UserId}
	update := bson.M{
		"$set": bson.M{
			"refresh_token":            params.RefreshToken,
			"refresh_token_expires_at": params.RefreshTokenExpiresAt,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := r.mongodb.Collection(User{}.TableName()).
		FindOneAndUpdate(ctx, filter, update, opts).
		Err()

	if !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	return nil
}
