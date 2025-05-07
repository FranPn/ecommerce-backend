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

	// User routes
	protected.GET("/profile", controllers.Profile)

	// Public product access
	protected.GET("/products", controllers.GetAllProducts)
	protected.GET("/products/:id", controllers.GetProductByID)

	// Returns order history for authenticated user
	protected.GET("/orders", controllers.GetUserOrders)

	// Creates a new order with selected products for the logged-in user
	protected.POST("/orders", controllers.CreateOrder)

	// Admin-only routes
	admin := protected.Group("/")
	admin.Use(middleware.AdminMiddleware())
	admin.POST("/products", controllers.CreateProduct)
	admin.PUT("/products/:id", controllers.UpdateProduct)
	admin.DELETE("/products/:id", controllers.DeleteProduct)
}
