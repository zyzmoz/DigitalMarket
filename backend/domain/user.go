package domain

import "gorm.io/gorm"

type Users []User

type User struct {
	gorm.Model
	ID        string `gorm:"primaryKey" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Validated bool   `json:"validated"`
}
