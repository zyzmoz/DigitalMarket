package domain

import "gorm.io/gorm"

type Products []Product

type Product struct {
	gorm.Model
	ID        string  `gorm:"primaryKey" json:"id"`
	Name      string  `json:"name"`
	Version   string  `json:"version"`
	Publisher *int64  `json:"publisher"`
	Price     float32 `json:"price"`
}
