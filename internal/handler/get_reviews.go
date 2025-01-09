package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMovieReviews(c *gin.Context) {
	movieId := c.Param("id")
	if movieId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "movie id is required"})
		return
	}

	reviews, err := h.biz.GetMovieReviews(c.Request.Context(), movieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.FailResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse(reviews))
}
