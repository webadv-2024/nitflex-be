package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) GetRecommendations(ctx context.Context, movieId int32) ([]*repository.SimilarMovieObj, error) {
	recommendedMovies, err := b.repo.GetRecommendationsByMovieId(ctx, movieId)
	if err != nil {
		return nil, err
	}

	for _, movie := range recommendedMovies {
		movie.PosterPath = "https://image.tmdb.org/t/p/w500" + movie.PosterPath
		movie.BackdropPath = "https://image.tmdb.org/t/p/w500" + movie.BackdropPath
	}

	return recommendedMovies, nil

}
