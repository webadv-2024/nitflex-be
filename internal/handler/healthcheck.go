package handler

import (
	"net/http"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HealthCheck(c *gin.Context) {

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse("OK"))
}