package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) GetUserByID(ctx context.Context, userID string) (*User, error) {
	var user User
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	err = r.mongodb.Collection(User{}.TableName()).
		FindOne(ctx, bson.M{"_id": objectID}).
		Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("user not found with id: %s", userID)
	}
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	return &user, nil
}

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

func (r *repository) UpdateUser(ctx context.Context, user *User) error {
	collection := r.mongodb.Collection(User{}.TableName())
	fmt.Println(user)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": user.Id}, bson.M{"$set": user})
	return err
}
