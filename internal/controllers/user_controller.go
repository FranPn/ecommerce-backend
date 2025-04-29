package controllers

import (
	"ecommerce-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register handles the user registration endpoint
func Register(c *gin.Context) {
	var input services.RegisterInput

	// Bind the JSON payload to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create a new user
	user, err := services.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Login handles the user login endpoint
func Login(c *gin.Context) {
	var input services.LoginInput

	// Bind the JSON payload to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to authenticate user
	token, err := services.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
