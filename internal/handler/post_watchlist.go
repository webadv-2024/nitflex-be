package handler

import (
	"errors"
	"net/http"
	"nitflex/constant"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

// WatchlistRequest represents the request body structure
type WatchlistRequest struct {
	MovieID string `json:"movie_id" binding:"required"`
}

func (h *Handler) PostWatchlist(c *gin.Context) {
	var (
		user = c.MustGet("user").(*util.JwtClaims)
		req    WatchlistRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	response, err := h.biz.UpdateWatchlist(c.Request.Context(), user.Id, req.MovieID)
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