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

	// Protected routes (requires JWT)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	// User routes
	protected.GET("/profile", controllers.Profile)

	// Product access (view)
	protected.GET("/products", controllers.GetAllProducts)
	protected.GET("/products/:id", controllers.GetProductByID)

	// Order routes for authenticated users
	protected.GET("/orders", controllers.GetUserOrders) // View order history
	protected.POST("/orders", controllers.CreateOrder)  // Create new order from input

	// Cart routes for authenticated users
	protected.POST("/cart", controllers.AddToCart)                    // Add or update item
	protected.GET("/cart", controllers.GetCart)                       // View cart
	protected.DELETE("/cart/:product_id", controllers.RemoveFromCart) // Remove item
	protected.POST("/cart/checkout", controllers.CheckoutCart)        // Checkout the cart and create an order

	// Admin-only product routes
	admin := protected.Group("/")
	admin.Use(middleware.AdminMiddleware())
	admin.POST("/products", controllers.CreateProduct)
	admin.PUT("/products/:id", controllers.UpdateProduct)
	admin.DELETE("/products/:id", controllers.DeleteProduct)
}
