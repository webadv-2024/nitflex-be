package handler

import (
	"net/http"
	"nitflex/internal/repository"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateReview(c *gin.Context) {
	// Get user ID from auth middleware
	user := c.MustGet("user").(*util.JwtClaims)
	userId := user.Id

	if userId == "" {
		c.JSON(http.StatusUnauthorized, util.FailResponse("unauthorized"))
		return
	}

	var req repository.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	err := h.biz.CreateMovieReview(c.Request.Context(), userId, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.FailResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse("Review created successfully"))
} 