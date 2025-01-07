package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *Handler) GetRecommendations(c *gin.Context) {
	var (
		movieId = c.Param("movie_id")
	)

	result, err := h.biz.GetRecommendations(c.Request.Context(), cast.ToInt32(movieId))
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
