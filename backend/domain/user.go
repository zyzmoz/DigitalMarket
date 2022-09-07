package domain

import "github.com/jinzhu/gorm"

type Users []User

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Validated bool   `json:"validated"`
}
