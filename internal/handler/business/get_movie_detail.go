package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) GetMovieDetail(ctx context.Context, id string) (*repository.Movie, error) {
	movie, err := b.repo.GetMovieByID(ctx, id)
	if err != nil {
		return nil, err
	}

	movie.PosterPath = "https://image.tmdb.org/t/p/w500" + movie.PosterPath
	movie.BackdropPath = "https://image.tmdb.org/t/p/w500" + movie.BackdropPath

	for i := range movie.Credits.Cast {
		movie.Credits.Cast[i].ProfilePath = "https://image.tmdb.org/t/p/w500" + movie.Credits.Cast[i].ProfilePath
	}

	return movie, nil
}
