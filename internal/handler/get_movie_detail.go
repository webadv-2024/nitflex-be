package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *Handler) GetMovieDetail(c *gin.Context) {
	var (
		id = c.Param("id")
	)

	result, err := h.tmdb.GetMovieInfo(cast.ToInt(id), nil)
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
