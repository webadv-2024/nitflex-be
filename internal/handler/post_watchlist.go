package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

// WatchlistRequest represents the request body structure
type WatchlistRequest struct {
	MovieID string `json:"movie_id" binding:"required"`
}

func (h *Handler) PostWatchlist(c *gin.Context) {
	var (
		userID = c.GetString("user_id")
		req    WatchlistRequest
	)

	if userID == "" {
		c.JSON(http.StatusUnauthorized, util.FailResponse("unauthorized"))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	// TODO: Save movie to user's watchlist in database
	// You'll need to implement this functionality in your business layer

	c.JSON(http.StatusOK, util.SuccessResponse(map[string]string{
		"user_id":  userID,
		"movie_id": req.MovieID,
	}))
}
