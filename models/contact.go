package models

import (
	"github.com/jinzhu/gorm"
)

// Contact contact model
type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserID uint `json:"user_id"`
}