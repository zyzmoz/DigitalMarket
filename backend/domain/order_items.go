package domain

import "gorm.io/gorm"

type OrderItems []OrderItem

type OrderItem struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	Product string
	Price   float32
}
