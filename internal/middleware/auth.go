package middleware

import (
	"net/http"
	"nitflex/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, util.FailResponse("authorization header is required"))
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, util.FailResponse("invalid authorization header format"))
			c.Abort()
			return
		}

		// Validate the token
		claims, err := util.Verify(c, parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, util.FailResponse("invalid token"))
			c.Abort()
			return
		}

		// Store user information in the context
		c.Set("user", claims)
		c.Next()
	}
} 