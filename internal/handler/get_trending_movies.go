package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"nitflex/util"
)

func (h *Handler) GetTrendingMovies(c *gin.Context) {
	var (
		timeWindow = c.Query("time_window")
	)

	// run biz
	response, err := h.biz.GetTrendingMovies(context.Background(), timeWindow)
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(response))
}
