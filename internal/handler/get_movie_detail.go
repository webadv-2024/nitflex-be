package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"nitflex/util"
	"strconv"
)

func (h *Handler) GetMovieDetail(c *gin.Context) {
	var (
		id = c.Param("id")
	)

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil || id == "" {
		c.JSON(http.StatusOK, util.FailResponse("Invalid id"))
		return
	}

	// run biz
	response, err := h.biz.GetMovieDetail(context.Background(), int(idInt))
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(response))
}
