package business

import (
	"context"
	"errors"
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

	// Check if movie is already in watchlist
	for _, id := range user.Watchlist {
		if id == movie.TmdbId {
			return &models.UpdateWatchlistResponse{
				Message: "Movie already in watchlist",
			}, nil
		}
	}

	user.Watchlist = append(user.Watchlist, movie.TmdbId)
	user.UpdatedAt = time.Now()

	err = b.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	return &models.UpdateWatchlistResponse{
		Message: "Movie added to watchlist",
	}, nil
}
