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

func (r *repository) GetMoviesList(ctx context.Context, movieIDs []int) ([]*Movie, error) {
	var movies []*Movie
	
	collection := r.mongodb.Collection(Movie{}.TableName())
	
	// Create filter for multiple movie IDs
	filter := bson.M{"tmdb_id": bson.M{"$in": movieIDs}}
	
	// Find all matching movies
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding movies: %v", err)
	}
	defer cursor.Close(ctx)
	
	// Decode all movies into the slice
	if err := cursor.All(ctx, &movies); err != nil {
		return nil, fmt.Errorf("error decoding movies: %v", err)
	}
	
	return movies, nil
}
