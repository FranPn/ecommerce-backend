package models

import (
	"time"

	"gorm.io/gorm"
)

// Product represents a product item in the store.
type Product struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
