package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
