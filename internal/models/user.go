package models

import "gorm.io/gorm"

// User represents the user model stored in the database
type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}
