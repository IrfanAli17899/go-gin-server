package migrations

import (
	"go-gin-server/config"
	"go-gin-server/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{})
}
