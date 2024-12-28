package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
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
		page       = c.Query("page")
	)

	var (
		result any
		err    error
	)

	if page == "" {
		page = "1"
	}

	if min_rating == "" {
		min_rating = "0"
	}

	if max_rating == "" {
		max_rating = "10"
	}

	if query != "" {
		result, err = h.tmdb.SearchMovie(query, map[string]string{
			"page": page,
		})
	} else {
		result, err = h.tmdb.DiscoverMovie(map[string]string{
			"with_genres":              genres,
			"vote_average.gte":         min_rating,
			"vote_average.lte":         max_rating,
			"primary_release_date.gte": from_date,
			"primary_release_date.lte": to_date,
			"with_cast":                actors,
			"page":                     page,
		})
	}

	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response.PosterPath = "https://image.tmdb.org/t/p/w500" + response.PosterPath
	// response.BackdropPath = "https://image.tmdb.org/t/p/w500" + response.BackdropPath

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
