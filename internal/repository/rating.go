package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) GetRatingUser(ctx context.Context, userID string) ([]*Rating, error) {
	// Find all ratings for the user
	var ratings []*Rating
	filter := bson.M{"user_id": userID}
	
	cursor, err := r.mongodb.Collection("ratings").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &ratings); err != nil {
		return nil, err
	}

	if len(ratings) == 0 {
		return make([]*Rating, 0), nil // Return empty slice if no ratings found
	}
	return ratings, nil
}

func (r *repository) CreateRating(ctx context.Context, userID string, movieID string, rating int) (*Rating, error) {
	// Validate rating range
	if rating < 1 || rating > 10 {
		return nil, fmt.Errorf("rating must be between 1 and 10")
	}

	// Find existing rating
	var existingRating Rating
	filter := bson.M{"user_id": userID, "movie_id": movieID}
	
	err := r.mongodb.Collection("ratings").FindOne(ctx, filter).Decode(&existingRating)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Create new rating if not exists
			newRating := Rating{
				UserId:  userID,
				MovieId: movieID,
				Rating:   rating,
			}
			
			_, err := r.mongodb.Collection("ratings").InsertOne(ctx, newRating)
			if err != nil {
				return nil, err
			}
			
			return &newRating, nil
		}
		return nil, err
	}

	// Update existing rating
	update := bson.M{"$set": bson.M{"rating": rating}}
	_, err = r.mongodb.Collection("ratings").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	existingRating.Rating = rating
	return &existingRating, nil
}