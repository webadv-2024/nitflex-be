package models

import "nitflex/internal/repository"

type MovieRating struct {
	ID          int    `bson:"id"`
	Title       string `bson:"title"`
	Overview    string `bson:"overview"`
	PosterPath  string `bson:"poster_path"`
}

type RatingWithMovie struct {
	*repository.Rating
	Movie *MovieRating `bson:"movie"`
}

type RatingsResponse struct {
	Results []*RatingWithMovie `bson:"results"`
}
