package domain

import "gorm.io/gorm"

type Products []Product

type Product struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Name      string
	Version   string
	Publisher *int64
	Price     float32
}
