package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nitflex/internal/models"
	"nitflex/util"
)

func (h *Handler) Register(c *gin.Context) {
	// parse request
	var request models.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	// run biz
	err := h.biz.Register(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(nil))
}
