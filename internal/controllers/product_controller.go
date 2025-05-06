package controllers

import (
	"ecommerce-backend/internal/models"
	"ecommerce-backend/pkg/db"
	"ecommerce-backend/pkg/search"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProduct handles POST /products
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetAllProducts handles GET /products
func GetAllProducts(c *gin.Context) {
	var products []models.Product

	if err := db.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductByID handles GET /products/:id
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles PUT /products/:id
// UpdateProduct handles PUT /products/:id
// It updates a product in the database and Meilisearch
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Check if product exists
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Bind JSON input
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.ImageURL = input.ImageURL

	// Save to DB
	if err := db.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Update in Meilisearch index
	go func() {
		search.Client.Index("products").UpdateDocuments([]interface{}{product})
	}()

	c.JSON(http.StatusOK, product)
}

// DeleteProduct handles DELETE /products/:id
// DeleteProduct handles DELETE /products/:id
// It deletes a product from the database and Meilisearch
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Check if product exists
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Delete from DB
	if err := db.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	// Remove from Meilisearch index
	go func() {
		search.Client.Index("products").DeleteDocument(strconv.Itoa(int(product.ID)))
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
