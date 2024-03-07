package domain

import "gorm.io/gorm"

type Orders []Order

type Order struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	User    *int64
	Status  string
	Price   float32
}
