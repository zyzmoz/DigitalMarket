package domain

import "github.com/jinzhu/gorm"

type Payments []Payment

type Payment struct {
	gorm.Model
	ProviderVerification string  `json:"providerVerification"`
	Product              int     `json:"product"`
	Total                float32 `json:"total"`
}
