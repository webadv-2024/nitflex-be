package handler

import (
	"github.com/cyruzin/go-tmdb"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	adapter "nitflex/internal/adapter/tmdb"
	"nitflex/internal/handler/business"
)

type Handler struct {
	biz  business.Business
	tmdb *tmdb.TMDb
}

func NewHandler(gormDb *gorm.DB, mongodb *mongo.Database) Handler {
	tmAdapter := adapter.NewTmdbAdapter()
	return Handler{
		biz: business.NewBusiness(gormDb, mongodb, tmAdapter),
		tmdb: tmdb.Init(tmdb.Config{
			APIKey: "565f77c9806ec9fab28d8de2a5257728",
		}),
	}
}
