package business

import (
	"context"
	"errors"
	"fmt"
	"time"

	"nitflex/constant"
	"nitflex/internal/models"
	"nitflex/internal/repository"
	"nitflex/util"

	"go.mongodb.org/mongo-driver/mongo"
)

func (b *business) UpdateWatchlist(ctx context.Context, userID string, movieID string) (*models.UpdateWatchlistResponse, error) {
	var (
		user *repository.User
		err  error
	)
	user, err = b.repo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, util.NewError(constant.ErrorMessage_NotFound)
		}

		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	movie, err := b.repo.GetMovieByID(ctx, movieID)

	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_NotFound)
	}

	user.Watchlist = append(user.Watchlist, movie.TmdbId)
	fmt.Println("//////")
	user.UpdatedAt = time.Now()

	err = b.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	return &models.UpdateWatchlistResponse{
		Message: "Movie added to watchlist",
	}, nil
}
