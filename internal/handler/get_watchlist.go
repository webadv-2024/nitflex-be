package handler

import (
	"errors"
	"net/http"
	"nitflex/constant"
	"nitflex/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWatchlist(c *gin.Context) {
	var (
		userID = c.GetString("user_id")
	)

	if userID == "" {
		c.JSON(http.StatusUnauthorized, util.FailResponse("unauthorized"))
		return
	}

	response, err := h.biz.GetWatchlist(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, util.NewError(constant.ErrorMessage_NotFound)) {
			c.JSON(http.StatusNotFound, util.FailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, util.FailResponse(constant.ErrorMessage_InternalServerError))
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse(response))
}