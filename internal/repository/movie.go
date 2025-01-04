package repository

import (
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) GetMovieByID(ctx context.Context, movieID string) (*Movie, error) {
	var movie Movie
	
	// Convert movieID string to int
	movieIDInt, err := strconv.Atoi(movieID)
	if err != nil {
		return nil, fmt.Errorf("invalid movie ID format: %v", err)
	}

	err = r.mongodb.Collection(Movie{}.TableName()).
		FindOne(ctx, bson.M{"tmdb_id": movieIDInt}).
		Decode(&movie)
	
	return &movie, err
}
