package domain

import "gorm.io/gorm"

type Orders []Order

type Order struct {
	gorm.Model
	ID      string  `gorm:"primaryKey" json:"id"`
	Product *int64  `json:"product"`
	User    *int64  `json:"user"`
	Price   float32 `json:"price"`
}
