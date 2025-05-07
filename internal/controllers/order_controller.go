package controllers

import (
	"ecommerce-backend/internal/models"
	"ecommerce-backend/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderInput struct {
	Items []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

// CreateOrder handles POST /orders
func CreateOrder(c *gin.Context) {
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract user ID from JWT claims (set by middleware)
	userID := c.GetUint("user_id")

	var total float64
	var orderItems []models.OrderItem

	// Loop through items to build order
	for _, item := range input.Items {
		var product models.Product
		if err := db.DB.First(&product, item.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
			return
		}

		subtotal := float64(item.Quantity) * product.Price
		total += subtotal

		orderItems = append(orderItems, models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
	}

	// Create and save order
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

	c.JSON(http.StatusCreated, order)
}

// GetUserOrders handles GET /orders
// Returns all orders for the logged-in user
func GetUserOrders(c *gin.Context) {
	userID := c.GetUint("user_id")

	var orders []models.Order
	err := db.DB.Preload("Items").Preload("Items.Product").Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
