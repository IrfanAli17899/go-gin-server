package controllers

import (
	"go-gin-server/models"
	"go-gin-server/services"
	"go-gin-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) (interface{}, *utils.AppError) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return nil, utils.NewAppError(http.StatusBadRequest, "Invalid request payload")
	}
	if err := services.CreateUser(&user); err != nil {
		return nil, utils.NewAppError(http.StatusInternalServerError, "Failed to create user")
	}
	return gin.H{"message": "User created successfully", "user": user}, nil
}
