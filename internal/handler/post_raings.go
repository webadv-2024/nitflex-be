package handler

import (
	"errors"
	"net/http"
	"nitflex/constant"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

// WatchlistRequest represents the request body structure
type RatingRequest struct {
	MovieID string `json:"movie_id" binding:"required"`
	Rating  int    `json:"rating" binding:"required"`
}

func (h *Handler) PostRating(c *gin.Context) {
	var (
		user = c.MustGet("user").(*util.JwtClaims)
		req    RatingRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	response, err := h.biz.UpdateRating(c.Request.Context(), user.Id, req.MovieID, req.Rating)
	if err != nil {
		if errors.Is(err, util.NewError(constant.ErrorMessage_NotFound)) {
			c.JSON(http.StatusNotFound, util.FailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, util.FailResponse(constant.ErrorMessage_InternalServerError))
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse(response))
}