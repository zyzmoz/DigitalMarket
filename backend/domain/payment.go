package domain

import "gorm.io/gorm"

type Payments []Payment

type Payment struct {
	gorm.Model
	ID                   string   `gorm:"primaryKey" json:"id"`
	Completed            bool     `json:"completed"`
	Product              *int     `json:"product"`
	ProviderVerification string   `json:"providerVerification"`
	Total                *float32 `json:"total"`
}
