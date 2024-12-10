package business

import (
	"context"
	"nitflex/constant"
	adapter "nitflex/internal/adapter/tmdb"
	"nitflex/internal/handler/models"
	"nitflex/util"
)

func (b *business) GetMovies(ctx context.Context, query string, page int) (*models.GetMoviesResponse, error) {
	if page <= 0 {
		page = 1
	}

	response, err := b.tmdbAdapter.SearchMovies(ctx, &adapter.SearchMoviesRequest{
		Query:    query,
		Language: constant.EN_US_Language,
		Page:     page,
	})
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	for _, movie := range response.Results {
		movie.PosterPath = "https://image.tmdb.org/t/p/w500" + movie.PosterPath
		movie.BackdropPath = "https://image.tmdb.org/t/p/w500" + movie.BackdropPath
	}

	return response, nil
}
