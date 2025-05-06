package main

import (
	"ecommerce-backend/internal/config"
	"ecommerce-backend/internal/routes"
	"ecommerce-backend/pkg/db"
	"ecommerce-backend/pkg/search"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	db.ConnectDatabase()

	// Initialize Meilisearch
	search.Init()

	// Create a new Gin router
	router := gin.Default()

	// Setup all application routes
	routes.SetupRoutes(router)

	// Start the server on port defined in .env
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
