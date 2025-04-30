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

	// User
	protected.GET("/profile", controllers.Profile)

	// Product
	protected.POST("/products", controllers.CreateProduct)
	protected.GET("/products", controllers.GetAllProducts)
	protected.GET("/products/:id", controllers.GetProductByID)
	protected.PUT("/products/:id", controllers.UpdateProduct)
	protected.DELETE("/products/:id", controllers.DeleteProduct)

	// Admin-only product routes
	admin := protected.Group("/")
	admin.Use(middleware.AdminMiddleware())
	admin.POST("/products", controllers.CreateProduct)
	admin.PUT("/products/:id", controllers.UpdateProduct)
	admin.DELETE("/products/:id", controllers.DeleteProduct)
}
