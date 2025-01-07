package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) GetGenres(ctx context.Context) ([]*Genre, error) {
	var genres []*Genre

	collection := r.mongodb.Collection("movie_genres")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &genres); err != nil {
		return nil, err
	}

	return genres, nil
}
