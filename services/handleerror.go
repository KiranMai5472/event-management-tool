package services

import (
	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/gin-gonic/gin"
)

// HandleError is used to handle the errors
func HandleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		Constants.Status:  Constants.Failed,
		Constants.Message: message,
		Constants.Code:    statusCode,
	})
}
