package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMe(c *gin.Context) {
	var (
		user = c.MustGet("user").(*util.JwtClaims)
	)

	result, err := h.biz.GetUserByUsername(c.Request.Context(), user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(result))
}
