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

	// Movies
	GetTrendingMovies(ctx context.Context, timeWindow string) ([]*repository.Movie, error)
	SearchMovies(ctx context.Context, query string) ([]*repository.Movie, error)
	FilterMovies(ctx context.Context, params *repository.FilterMoviesParams) ([]*repository.Movie, error)
	GetMovieDetail(ctx context.Context, id string) (*repository.Movie, error)

	// Cast
	GetCastDetail(ctx context.Context, tmdbid int32) (*repository.CastInfo, error)

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
	GetRatingUser(ctx context.Context, userID string) (*models.RatingsResponse, error)
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
