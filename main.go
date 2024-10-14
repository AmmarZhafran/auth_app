package main

import (
	"auth-app/database"
	"auth-app/models"
	"auth-app/routes"
	"log"
)

func main() {
	database.Connect()
	db := database.DB

	// Migrate database
	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes()

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
