package handler

import (
	"net/http"
	"nitflex/internal/handler/models"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GoogleLogin(c *gin.Context) {
	var request models.GoogleLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	response, err := h.biz.GoogleLogin(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse(response))
}
