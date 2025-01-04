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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (b *business) UpdateFavoriteList(ctx context.Context, userID string, movieID string) (*models.UpdateFavoriteListResponse, error) {
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

	// Check if movie is already in favorite list
	for _, id := range user.FavoriteList {
		if id == movie.TmdbId {
			return &models.UpdateFavoriteListResponse{
				Message: "Movie already in favorite list",
			}, nil
		}
	}

	user.FavoriteList = append(user.FavoriteList, movie.TmdbId)
	user.UpdatedAt = time.Now()

	err = b.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	return &models.UpdateFavoriteListResponse{
		Message: "Movie added to favorite list",
	}, nil
}

func (b *business) GetFavoriteList(ctx context.Context, userID string) (*models.GetFavoriteListResponse, error) {
	user, err := b.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_NotFound)
	}

	movies, err := b.repo.GetMoviesList(ctx, user.FavoriteList)
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

	return &models.GetFavoriteListResponse{
		Results: results,
	}, nil
}

func (b *business) RemoveFromFavoriteList(ctx context.Context, userID string, movieID string) (*models.UpdateFavoriteListResponse, error) {
	user, err := b.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_NotFound)
	}

	movieIDInt, err := strconv.Atoi(movieID)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_BadRequest)
	}
	user.FavoriteList = util.RemoveElement(user.FavoriteList, movieIDInt)

	err = b.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	return &models.UpdateFavoriteListResponse{
		Message: "Movie removed from favorite list",
	}, nil
}

func GetFavoriteList(db *mongo.Database, userID string) ([]map[string]interface{}, error) {
	collection := db.Collection("favorite_list")
	
	filter := bson.M{"user_id": userID}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []map[string]interface{}
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func DeleteFavoriteList(db *mongo.Database, userID string, movieID string) error {
	collection := db.Collection("favorite_list")
	
	filter := bson.M{
		"user_id": userID,
		"movie_id": movieID,
	}
	
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}