package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetGenres(c *gin.Context) {
	result, err := h.biz.GetGenres(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
