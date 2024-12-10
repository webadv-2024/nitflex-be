package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"nitflex/util"
	"strconv"
)

func (h *Handler) GetMovies(c *gin.Context) {
	var (
		query = c.Query("query")
		page  = c.Query("page")
	)

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil || page == "" {
		pageInt = 1
	}

	// run biz
	response, err := h.biz.GetMovies(context.Background(), query, int(pageInt))
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(response))
}
