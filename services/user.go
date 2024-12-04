package services

import (
	"go-gin-server/config"
	"go-gin-server/models"
)

func CreateUser(user *models.User) error {
	result := config.DB.Create(user)
	return result.Error
}
