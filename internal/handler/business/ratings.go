package business

import (
	"context"
	"nitflex/internal/models"
)

func (b *business) UpdateRating(ctx context.Context, userID string, movieID string, rating int) (*models.TextResponse, error) {
	_, err := b.repo.CreateRating(ctx, userID, movieID, rating)
	if err != nil {
		return nil, err
	}

	return &models.TextResponse{
		Message: "Rating updated successfully",
	}, nil
}

func (b *business) GetRatingUser(ctx context.Context, userID string) (*models.RatingsResponse, error) {
	ratings, err := b.repo.GetRatingUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	results := make([]*models.RatingWithMovie, 0, len(ratings))
	for _, r := range ratings {
		movie, err := b.repo.GetMovieByID(ctx, r.MovieId)
		if err != nil {
			return nil, err
		}
	
		ratingWithMovie := &models.RatingWithMovie{
			Rating: r,
			Movie: &models.MovieRating{
				ID:          movie.TmdbId,
				Title:       movie.Title,
				Overview:    movie.Overview,
				PosterPath:  movie.PosterPath,
			},
		}
		results = append(results, ratingWithMovie)
	}

	return &models.RatingsResponse{
		Results: results,
	}, nil
}
