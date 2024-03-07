package domain

import "gorm.io/gorm"

type Payments []Payment

type Payment struct {
	gorm.Model
	ID                   string `gorm:"primaryKey"`
	Order                string
	Completed            bool
	ProviderVerification string
	Total                *float32
}
