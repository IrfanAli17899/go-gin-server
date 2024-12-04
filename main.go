package main

import (
	"go-gin-server/config"
	"go-gin-server/middlewares"
	"go-gin-server/migrations"
	"go-gin-server/routes"
	"log"
)

func main() {
	if err := config.InitDB(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	migrations.Migrate()

	router := routes.Setup()
	router.Use(middlewares.ErrorHandlingMiddleware())
	router.Run(":8080")
}
