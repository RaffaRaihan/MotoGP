package main

import (
	"motogp-api/config"
	"motogp-api/models"
	"motogp-api/routes"
)

func main() {
	config.LoadEnv()
	config.Connect()

	// Migrasi model ke database
	config.GetDB().AutoMigrate(&models.Tim{}, &models.Pembalap{})

	router := routes.SetupRouter()
	router.Run(":8080")
}
