package business

import (
	"context"
	"nitflex/constant"
	adapter "nitflex/internal/adapter/tmdb"
	"nitflex/internal/handler/models"
	"nitflex/util"
)

func (b *business) GetMovieDetail(ctx context.Context, id int) (*models.GetMovieDetailResponse, error) {
	response, err := b.tmdbAdapter.GetMovieDetail(ctx, &adapter.GetMovieDetailRequest{
		Id: id,
	})
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}
	response.PosterPath = "https://image.tmdb.org/t/p/w500" + response.PosterPath
	response.BackdropPath = "https://image.tmdb.org/t/p/w500" + response.BackdropPath
	
	return response, nil
}
