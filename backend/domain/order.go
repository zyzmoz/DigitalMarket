package domain

import "gorm.io/gorm"

type Orders []Order

type Order struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	Product *int64
	User    *int64
	Price   float32
}
