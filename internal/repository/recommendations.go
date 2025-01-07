package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) GetRecommendationsByMovieId(ctx context.Context, tmdb_id int32) ([]*SimilarMovieObj, error) {
	var (
		similarMovie *SimilarMovie
		movies       []*SimilarMovieObj
		collection   = r.mongodb.Collection("similar")
		filter       = bson.M{"tmdb_id": tmdb_id}
	)

	// Find the cast by ID
	err := collection.FindOne(ctx, filter).Decode(&similarMovie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return movies, nil
		}
		return nil, fmt.Errorf("error finding cast: %v", err)
	}

	return similarMovie.SimilarMovie, nil
}
