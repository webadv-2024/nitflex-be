package handler

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"nitflex/internal/models"
	"nitflex/util"
)

func (h *Handler) ResetPassword(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Token is required"})
		return
	}

	// parse request
	var request models.ResetPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	// Xác thực token
	user, err := h.biz.VerifyResetPasswordToken(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
		return
	}

	// Cập nhật mật khẩu mới
	err = h.biz.UpdatePassword(c.Request.Context(), user.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update new password"})
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse(struct {
		Message string `json:"message"`
	}{
		Message: "Update password successfully",
	}))
}
