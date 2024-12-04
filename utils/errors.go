package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppError represents a custom error structure
type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// NewAppError creates a new AppError
func NewAppError(status int, message string) *AppError {
	return &AppError{
		Status:  status,
		Message: message,
	}
}

// Error implements the error interface for AppError
func (e *AppError) Error() string {
	return e.Message
}

// HandleError is a utility function to handle errors
func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Status, gin.H{"error": appErr.Message})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
	}
}
