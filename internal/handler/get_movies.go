package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"nitflex/internal/repository"
	"nitflex/util"
)

func (h *Handler) GetMovies(c *gin.Context) {
	var (
		query      = c.Query("query")
		genres     = c.Query("genres")
		min_rating = c.Query("min_rating")
		max_rating = c.Query("max_rating")
		from_date  = c.Query("release_date_gte") //YYYY-MM-DD
		to_date    = c.Query("release_date_lte") //YYYY-MM-DD
		actors     = c.Query("actors")
		// page       = c.Query("page")
	)

	var (
		result = []*repository.Movie{}
		err    error
		ctx    = context.Background()
	)

	if query != "" {
		result, err = h.biz.SearchMovies(ctx, query)
	} else {
		result, err = h.biz.FilterMovies(ctx, &repository.FilterMoviesParams{
			Genres:         genres,
			MinRating:      cast.ToFloat64(min_rating),
			MaxRating:      cast.ToFloat64(max_rating),
			ReleaseDateGte: from_date,
			ReleaseDateLte: to_date,
			Actors:         actors,
		})
	}

	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
