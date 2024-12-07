package business

import (
	"context"
	models2 "nitflex/internal/handler/models"
	"nitflex/internal/repository"

	"gorm.io/gorm"
)

type Business interface {
	Register(ctx context.Context, request *models2.RegisterRequest) error
	Login(ctx context.Context, request *models2.LoginRequest) (*models2.LoginResponse, error)
	GoogleLogin(ctx context.Context, request *models2.GoogleLoginRequest) (*models2.LoginResponse, error)
}

type business struct {
	repo repository.Repository
}

func NewBusiness(gormDb *gorm.DB) Business {
	return &business{
		repo: repository.NewRepository(gormDb),
	}
}
