package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) GetMoviesByIDs(ctx context.Context, movieIDs []string) ([]*repository.Movie, error) {
	movies, err := b.repo.GetMoviesListByObjectIds(ctx, movieIDs)
	if err != nil {
		return nil, err
	}

	for i := range movies {
		movies[i].PosterPath = "https://image.tmdb.org/t/p/w500" + movies[i].PosterPath
		movies[i].BackdropPath = "https://image.tmdb.org/t/p/w500" + movies[i].BackdropPath
	}

	return movies, nil
}
