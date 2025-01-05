package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) CreateRating(ctx context.Context, userID string, movieID string, rating int) (*Rating, error) {
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