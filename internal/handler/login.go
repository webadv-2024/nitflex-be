package handler

import (
	"net/http"
	"nitflex/internal/handler/models"

	"github.com/gin-gonic/gin"
	"nitflex/util"
)

func (h *Handler) Login(c *gin.Context) {
	// parse request
	var request models.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	// run biz
	response, err := h.biz.Login(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(response))
}
