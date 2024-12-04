package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ControllerFunc defines a function type for controller handlers
type ControllerFunc func(c *gin.Context) (interface{}, *AppError)

// WrapController wraps a controller function with validation, error handling, and response formatting
func WrapController(fn ControllerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate request body if needed
		if err := c.ShouldBind(&c.Request.Body); err != nil {
			HandleError(c, NewAppError(http.StatusBadRequest, "Invalid request payload"))
			return
		}

		// Handle the controller function
		data, appErr := fn(c)
		if appErr != nil {
			HandleError(c, appErr)
			return
		}

		// Respond with data
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
