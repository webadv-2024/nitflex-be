package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	fmt.Println("Find user with email: ", email)
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
	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %w", err)
	}

	// Prepare the fields to update (exclude the ID field)
	update := bson.M{
		"$set": bson.M{
			"username":      user.Username,
			"email":         user.Email,
			"password":      user.Password,
			"updated_at":    user.UpdatedAt, // Ensure you set this to the current timestamp before calling the function
			"watchlist":     user.Watchlist,
			"favorite_list": user.FavoriteList,
		},
	}

	// Access the collection
	collection := r.mongodb.Collection(User{}.TableName())

	// Perform the update
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (r *repository) GetUserByActivationToken(ctx context.Context, token string) (*User, error) {
	var user User
	err := r.mongodb.Collection(User{}.TableName()).FindOne(ctx, bson.M{"activation_token": token}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByResetPasswordToken(ctx context.Context, token string) (*User, error) {
	var user User
	err := r.mongodb.Collection(User{}.TableName()).FindOne(ctx, bson.M{"reset_password_token": token}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) UpdateUserActivationStatus(ctx context.Context, username string, isActive bool) error {
	_, err := r.mongodb.Collection(User{}.TableName()).UpdateOne(ctx, bson.M{"username": username}, bson.M{
		"$set": bson.M{"is_activated": isActive, "activation_token": ""},
	})
	return err
}

func (r *repository) UpdateResetPasswordToken(ctx context.Context, username string, token string, tokenExpiresAt time.Time) error {
	_, err := r.mongodb.Collection(User{}.TableName()).UpdateOne(ctx, bson.M{"username": username}, bson.M{
		"$set": bson.M{"reset_password_token": token, "reset_password_token_expires_at": tokenExpiresAt},
	})
	return err
}

func (r *repository) UpdatePassword(ctx context.Context, username string, password string) error {
	_, err := r.mongodb.Collection(User{}.TableName()).UpdateOne(ctx, bson.M{"username": username}, bson.M{
		"$set": bson.M{"password": password, "reset_password_token": "", "reset_password_token_expires_at": time.Time{}},
	})
	return err
}
