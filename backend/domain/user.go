package domain

import "gorm.io/gorm"

type Users []User

type User struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Validated bool
}
