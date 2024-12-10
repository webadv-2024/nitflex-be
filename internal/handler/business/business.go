package business

import (
	"context"
	adapter "nitflex/internal/adapter/tmdb"

	"gorm.io/gorm"

	models2 "nitflex/internal/handler/models"
	"nitflex/internal/repository"
)

type Business interface {
	Register(ctx context.Context, request *models2.RegisterRequest) error
	Login(ctx context.Context, request *models2.LoginRequest) (*models2.LoginResponse, error)
	GoogleLogin(ctx context.Context, request *models2.GoogleLoginRequest) (*models2.LoginResponse, error)

	GetTrendingMovies(ctx context.Context, timeWindow string) ([]*models2.Movie, error)
	GetMovies(ctx context.Context, query string, page int) (*models2.GetMoviesResponse, error)
}

type business struct {
	repo repository.Repository

	tmdbAdapter adapter.TmdbAdapter
}

func NewBusiness(gormDb *gorm.DB, tmdbAdapter adapter.TmdbAdapter) Business {
	return &business{
		repo:        repository.NewRepository(gormDb),
		tmdbAdapter: tmdbAdapter,
	}
}
