package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repository) GetReviewsByMovieId(ctx context.Context, movieId string) ([]*Review, error) {
	collection := r.mongodb.Collection("reviews")
	
	filter := bson.M{"movie_id": movieId}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []*Review
	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *repository) CreateReview(ctx context.Context, review *Review) error {
	collection := r.mongodb.Collection("reviews")
	
	review.Id = primitive.NewObjectID().Hex()
	review.CreatedAt = time.Now()
	
	_, err := collection.InsertOne(ctx, review)
	return err
}
