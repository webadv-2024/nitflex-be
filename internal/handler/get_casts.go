package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *Handler) GetCasts(c *gin.Context) {
	var (
		movieId = c.Query("movie_id")
	)

	result, err := h.tmdb.GetMovieCredits(cast.ToInt(movieId), nil)
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
