package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	// register routes
	UserRoutes(router)

	return router
}
