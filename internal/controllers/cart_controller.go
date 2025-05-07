package controllers

import (
	"ecommerce-backend/internal/models"
	"ecommerce-backend/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddToCartInput struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

// AddToCart adds or updates a product in the user's cart
func AddToCart(c *gin.Context) {
	var input AddToCartInput
	if err := c.ShouldBindJSON(&input); err != nil || input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetUint("user_id")
	var cartItem models.CartItem

	// Check if item already in cart
	if err := db.DB.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&cartItem).Error; err == nil {
		// Update quantity
		cartItem.Quantity += input.Quantity
		db.DB.Save(&cartItem)
	} else {
		// Create new cart item
		newItem := models.CartItem{
			UserID:    userID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		db.DB.Create(&newItem)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}

// GetCart retrieves all items in the user's cart
func GetCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	var cart []models.CartItem

	if err := db.DB.Preload("Product").Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

// RemoveFromCart deletes a specific product from the cart
func RemoveFromCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	productID := c.Param("product_id")

	if err := db.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}

// CheckoutCart creates an order from the user's cart and clears it
func CheckoutCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cartItems []models.CartItem
	if err := db.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	var orderItems []models.OrderItem
	var total float64

	for _, item := range cartItems {
		subtotal := float64(item.Quantity) * item.Product.Price
		total += subtotal

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})
	}

	order := models.Order{
		UserID: userID,
		Items:  orderItems,
		Total:  total,
		Status: "pending",
	}

	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Clear cart
	db.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})

	c.JSON(http.StatusCreated, order)
}
