package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID uint        `json:"user_id"`
	User   User        `gorm:"foreignKey:UserID"`
	Items  []OrderItem `json:"items"`
	Total  float64     `json:"total"`
	Status string      `json:"status"` // e.g. "pending", "completed"
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // snapshot of product price at time of order
}
