package business

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	adapter "nitflex/internal/adapter/tmdb"
	"nitflex/internal/models"
	"nitflex/internal/repository"
)

type Business interface {
	Register(ctx context.Context, request *models.RegisterRequest) error
	Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error)
	GoogleLogin(ctx context.Context, request *models.GoogleLoginRequest) (*models.LoginResponse, error)

	GetTrendingMovies(ctx context.Context, timeWindow string) ([]*models.Movie, error)

	// Watchlist methods
	UpdateWatchlist(ctx context.Context, userID string, movieID string) (*models.TextResponse, error)
	GetWatchlist(ctx context.Context, userID string) (*models.MovieListResponse, error)
	RemoveFromWatchlist(ctx context.Context, userID string, movieID string) (*models.TextResponse, error)

	// Favorite list methods
	UpdateFavoriteList(ctx context.Context, userID string, movieID string) (*models.TextResponse, error)
	GetFavoriteList(ctx context.Context, userID string) (*models.MovieListResponse, error)
	RemoveFromFavoriteList(ctx context.Context, userID string, movieID string) (*models.TextResponse, error)

	// Rating methods
	UpdateRating(ctx context.Context, userID string, movieID string, rating int) (*models.TextResponse, error)
}

type business struct {
	repo        repository.Repository
	tmdbAdapter adapter.TmdbAdapter
}

func NewBusiness(gormDb *gorm.DB, mongodb *mongo.Database, tmdbAdapter adapter.TmdbAdapter) Business {
	return &business{
		repo:        repository.NewRepository(gormDb, mongodb),
		tmdbAdapter: tmdbAdapter,
	}
}
