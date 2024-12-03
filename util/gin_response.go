package util

import "github.com/gin-gonic/gin"

func SuccessResponse(data any) gin.H {
	return gin.H{
		"success": true,
		"message": "",
		"data":    data,
	}
}

func FailResponse(message string) gin.H {
	return gin.H{
		"success": false,
		"message": message,
	}
}
