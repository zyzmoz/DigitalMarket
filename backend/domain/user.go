package domain

import "gorm.io/gorm"

type Users []User

type User struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Validated bool
}
