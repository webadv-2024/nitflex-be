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
