package handler

import (
	"os"

	"github.com/cyruzin/go-tmdb"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	"nitflex/internal/adapter/llm_search"
	adapter "nitflex/internal/adapter/tmdb"
	"nitflex/internal/handler/business"
)

type Handler struct {
	biz  business.Business
	tmdb *tmdb.TMDb
}

func NewHandler(gormDb *gorm.DB, mongodb *mongo.Database) Handler {
	tmAdapter := adapter.NewTmdbAdapter()
	cfg := llm_search.Config{
		OpenAIApiKey: os.Getenv("OPENAI_API_KEY"),
	}
	llmAdapter := llm_search.NewLLMSearchAdapter(cfg)
	return Handler{
		biz: business.NewBusiness(gormDb, mongodb, tmAdapter, llmAdapter),
		tmdb: tmdb.Init(tmdb.Config{
			APIKey: "565f77c9806ec9fab28d8de2a5257728",
		}),
	}
}
