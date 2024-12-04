package routes

import (
	"go-gin-server/controllers"
	"go-gin-server/utils"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.POST("/", utils.WrapController(controllers.CreateUser))
	}
}
