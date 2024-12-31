package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMovieUpcoming(c *gin.Context) {
	var (
		language = c.Query("language")
		page     = c.Query("page")
	)

	if language == "" {
		language = "en-US"
	}
	if page == "" {
		page = "1"
	}

	options := map[string]string{
		"language": language,
		"page":     page,
	}

	result, err := h.tmdb.GetMovieUpcoming(options)
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
