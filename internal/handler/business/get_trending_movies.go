package business

import (
	"context"
	"nitflex/util"

	"nitflex/constant"

	adapter "nitflex/internal/adapter/tmdb"
	"nitflex/internal/handler/models"
)

func (b *business) GetTrendingMovies(ctx context.Context, timeWindow string) ([]*models.Movie, error) {
	if timeWindow != constant.TrendingMovies_TimeWindow_Day && timeWindow != constant.TrendingMovies_TimeWindow_Week {
		timeWindow = constant.TrendingMovies_TimeWindow_Day
	}

	response, err := b.tmdbAdapter.GetTrendingMovies(ctx, &adapter.GetTrendingMoviesRequest{
		TimeWindow: timeWindow,
		Language:   constant.EN_US_Language,
	})
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	for _, movie := range response.Results {
		movie.PosterPath = "https://image.tmdb.org/t/p/w500" + movie.PosterPath
		movie.BackdropPath = "https://image.tmdb.org/t/p/w500" + movie.BackdropPath
	}

	return response.Results, nil
}
