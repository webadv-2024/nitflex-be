package business

import (
	"context"
	"fmt"
	"nitflex/internal/repository"
)

func (b *business) GetMovieReviews(ctx context.Context, movieId string) ([]*repository.ReviewResponse, error) {
	reviews, err := b.repo.GetReviewsByMovieId(ctx, movieId)
	if err != nil {
		return nil, err
	}

	var response []*repository.ReviewResponse
	for _, review := range reviews {
		user, err := b.repo.GetUserByID(ctx, review.UserId)
		if err != nil {
			continue
		}

		response = append(response, &repository.ReviewResponse{
			Id:      review.Id,
			Author:  user.Username,
			Content: review.Content,
		})
	}

	return response, nil
}

func (b *business) CreateMovieReview(ctx context.Context, userId string, req *repository.CreateReviewRequest) error {
	// Validate if movie exists
	_, err := b.repo.GetMovieByIdObject(ctx, req.MovieId)
	if err != nil {
		return fmt.Errorf("movie not found")
	}

	review := &repository.Review{
		UserId:  userId,
		MovieId: req.MovieId,
		Content: req.Content,
	}

	return b.repo.CreateReview(ctx, review)
}
