package routes

import (
	"ecommerce-backend/internal/controllers"
	"ecommerce-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes defines the application's routes
func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", controllers.Profile)
}
