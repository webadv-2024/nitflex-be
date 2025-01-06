package business

import (
	"context"

	"nitflex/constant"
	"nitflex/internal/repository"
)

func (b *business) GetTrendingMovies(ctx context.Context, timeWindow string) ([]*repository.Movie, error) {
	var (
		result = []*repository.Movie{}
		err    error
	)

	if timeWindow == constant.TrendingMovies_TimeWindow_Week {

	} else {
		result, err = b.repo.GetTrendingMoviesInDay(ctx)
	}
	if err != nil {
		return nil, err
	}

	for _, movie := range result {
		movie.PosterPath = "https://image.tmdb.org/t/p/w500" + movie.PosterPath
		movie.BackdropPath = "https://image.tmdb.org/t/p/w500" + movie.BackdropPath
	}

	return result, nil
}
