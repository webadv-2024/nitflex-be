package business

import (
	"context"
	"errors"
	"strconv"
	"time"

	"nitflex/constant"
	"nitflex/internal/models"
	"nitflex/internal/repository"
	"nitflex/util"

	"go.mongodb.org/mongo-driver/mongo"
)

func (b *business) UpdateWatchlist(ctx context.Context, userID string, movieID string) (*models.TextResponse, error) {
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
			return &models.TextResponse{
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

	return &models.TextResponse{
		Message: "Movie added to watchlist",
	}, nil
}

func (b *business) GetWatchlist(ctx context.Context, userID string) (*models.MovieListResponse, error) {
	user, err := b.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_NotFound)
	}

	if len(user.Watchlist) == 0 {
		return &models.MovieListResponse{
			Results: []*models.Movie{},
		}, nil
	}

	// Get all movies in one database call
	movies, err := b.repo.GetMoviesList(ctx, user.Watchlist)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	results := make([]*models.Movie, len(movies))
	for i, movie := range movies {
		results[i] = &models.Movie{
			Id:            movie.TmdbId,
			Title:         movie.Title,
			OriginalTitle: movie.OriginalTitle,
			Overview:      movie.Overview,
			PosterPath:    movie.PosterPath,
			ReleaseDate:   movie.ReleaseDate,
			VoteAverage:   movie.VoteAverage,
			VoteCount:     movie.VoteCount,
		}
	}

	return &models.MovieListResponse{
		Results: results,
	}, nil
}

func (b *business) RemoveFromWatchlist(ctx context.Context, userID string, movieID string) (*models.TextResponse, error) {
	user, err := b.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_NotFound)
	}

	movieIDInt, err := strconv.Atoi(movieID)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_BadRequest)
	}
	user.Watchlist = util.RemoveElement(user.Watchlist, movieIDInt)

	err = b.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	return &models.TextResponse{
		Message: "Movie removed from watchlist",
	}, nil

}
