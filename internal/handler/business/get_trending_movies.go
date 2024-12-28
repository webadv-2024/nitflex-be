package business

import (
	"context"
	"nitflex/util"

	"nitflex/constant"

	"nitflex/internal/models"
)

func (b *business) GetTrendingMovies(ctx context.Context, timeWindow string) ([]*models.Movie, error) {
	if timeWindow != constant.TrendingMovies_TimeWindow_Day && timeWindow != constant.TrendingMovies_TimeWindow_Week {
		timeWindow = constant.TrendingMovies_TimeWindow_Day
	}

	response, err := b.tmdbAdapter.GetTrendingMovies(ctx, timeWindow)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	for _, movie := range response.Results {
		movie.PosterPath = "https://image.tmdb.org/t/p/w500" + movie.PosterPath
		movie.BackdropPath = "https://image.tmdb.org/t/p/w500" + movie.BackdropPath
	}

	return response.Results, nil
}
