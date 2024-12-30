package handler

import (
	"net/http"
	"nitflex/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetReviews(c *gin.Context) {
	var (
		idStr    = c.Param("id")
		language = c.Query("language")
		page     = c.Query("page")
	)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse("Invalid ID"))
		return
	}

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

	result, err := h.tmdb.GetMovieReviews(id, options)
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
