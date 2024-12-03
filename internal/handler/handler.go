package handler

import (
	"gorm.io/gorm"
	"nitflex/internal/handler/business"
)

type Handler struct {
	biz business.Business
}

func NewHandler(gormDb *gorm.DB) Handler {
	return Handler{
		biz: business.NewBusiness(gormDb),
	}
}
