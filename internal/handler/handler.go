package handler

import (
	"gorm.io/gorm"

	adapter "nitflex/internal/adapter/tmdb"
	"nitflex/internal/handler/business"
)

type Handler struct {
	biz business.Business
}

func NewHandler(gormDb *gorm.DB) Handler {
	tmAdapter := adapter.NewTmdbAdapter()
	return Handler{
		biz: business.NewBusiness(gormDb, tmAdapter),
	}
}
