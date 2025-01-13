package handler

import (
	"net/http"
	"nitflex/internal/models"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) RequestResetPassword(ctx *gin.Context) {
	// parse request
	var request models.RequestResetPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	// run biz
	err := handler.biz.RequestResetPassword(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	ctx.JSON(http.StatusOK, util.SuccessResponse(nil))
}
