package handler

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"nitflex/util"
)

func (h *Handler) ActivateEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Token is required"})
		return
	}

	// Xác thực token
	user, err := h.biz.VerifyActivationToken(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
		return
	}

	// Kích hoạt tài khoản
	err = h.biz.ActivateUser(c.Request.Context(), user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to activate account"})
		return
	}

	c.JSON(http.StatusOK, util.SuccessResponse(struct {
		Message string `json:"message"`
	}{
		Message: "Account activated successfully",
	}))
}
