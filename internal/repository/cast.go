package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) GetCastByID(ctx context.Context, tmdb_id int32) (*CastInfo, error) {
	var (
		castInfo   CastInfo
		collection = r.mongodb.Collection("people")
		filter     = bson.M{"tmdb_id": tmdb_id}
	)

	// Find the cast by ID
	err := collection.FindOne(ctx, filter).Decode(&castInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("cast not found")
		}
		return nil, fmt.Errorf("error finding cast: %v", err)
	}

	return &castInfo, nil
}
