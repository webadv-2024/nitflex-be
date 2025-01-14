package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

type LLMSearchRequest struct {
	Description string `json:"description"`
}

// GetLLMSearch handles movie search requests using LLM
func (h *Handler) PostLLMSearch(c *gin.Context) {
	response := gin.H{
		"data": []string{"test"},
	}

	c.JSON(http.StatusOK, util.SuccessResponse(response))
}
