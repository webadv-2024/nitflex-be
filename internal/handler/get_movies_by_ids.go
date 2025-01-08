package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"nitflex/internal/repository"
	"nitflex/util"
)

func (h *Handler) GetMoviesByIDs(c *gin.Context) {
	var (
		movieIDs = c.Query("ids")
	)

	fmt.Println(movieIDs)

	var (
		result = []*repository.Movie{}
		err    error
		ctx    = context.Background()
	)

	movieIDsSlice := strings.Split(movieIDs, ",")

	if movieIDs != "" {
		result, err = h.biz.GetMoviesByIDs(ctx, movieIDsSlice)
	} else {
		c.JSON(http.StatusOK, util.FailResponse("movie_ids is required"))
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
