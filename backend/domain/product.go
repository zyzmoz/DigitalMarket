package domain

import "gorm.io/gorm"

type Products []Product

type Product struct {
	gorm.Model
	Name      string  `json:"name"`
	Version   string  `json:"version"`
	Publisher *int64  `json:"publisher"`
	Price     float32 `json:"price"`
}
