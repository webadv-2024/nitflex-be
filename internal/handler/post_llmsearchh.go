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
	var req LLMSearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	if req.Description == "" {
		c.JSON(http.StatusBadRequest, util.FailResponse("Description is required"))
		return
	}

	movies, err := h.biz.SearchMoviesLLM(c.Request.Context(), req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.FailResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse(movies))
}
