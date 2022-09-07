package infra

import (
	"github.com/jinzhu/gorm"
	"github.com/zyzmoz/DigitalMarket/database"
	"github.com/zyzmoz/DigitalMarket/domain"
)

func InitDb() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "temp.db")

	if err != nil {
		panic("failed to create database connection")
	}

	database.DBConn.AutoMigrate(&domain.User{})
}
