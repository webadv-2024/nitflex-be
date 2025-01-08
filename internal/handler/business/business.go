package business

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	llm_search "nitflex/internal/adapter/llm_search"
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

	// Genres
	GetGenres(ctx context.Context) ([]*repository.Genre, error)

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

	// Recommendations
	GetRecommendations(ctx context.Context, movieId int32) ([]*repository.SimilarMovieObj, error)

	// LLM Search
	SearchMoviesLLM(ctx context.Context, description string) ([]*repository.Movie, error)

	// Get movies by ids
	GetMoviesByIDs(ctx context.Context, movieIDs []string) ([]*repository.Movie, error)
}

type business struct {
	repo        repository.Repository
	tmdbAdapter adapter.TmdbAdapter
	llmAdapter  llm_search.LLMSearchAdapter
}

func NewBusiness(gormDb *gorm.DB, mongodb *mongo.Database, tmdbAdapter adapter.TmdbAdapter, llmAdapter llm_search.LLMSearchAdapter) Business {
	return &business{
		repo:        repository.NewRepository(gormDb, mongodb),
		tmdbAdapter: tmdbAdapter,
		llmAdapter:  llmAdapter,
	}
}