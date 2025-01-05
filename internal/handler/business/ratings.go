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
	rating, err := b.repo.GetRatingUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &models.RatingsResponse{
		Results: rating,
	}, nil
}