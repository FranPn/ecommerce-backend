package services

import (
	"ecommerce-backend/internal/models"
	"ecommerce-backend/internal/utils"
	"ecommerce-backend/pkg/db"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Struct for registration input
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// Struct for login input
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterUser handles the creation of a new user
func RegisterUser(input RegisterInput) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

// LoginUser handles the user login logic
func LoginUser(input LoginInput) (string, error) {
	var user models.User

	result := db.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		return "", errors.New("invalid email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
