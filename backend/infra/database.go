package infra

import (
	"os"

	"github.com/zyzmoz/DigitalMarket/database"
	"github.com/zyzmoz/DigitalMarket/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDb() {
	var err error

	switch os.Getenv("DATABASE_DRIVER") {
	case "mssql":
		dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
		database.DBConn, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	default:
		database.DBConn, err = gorm.Open(sqlite.Open("temp.db"), &gorm.Config{})
	}

	if err != nil {
		panic("failed to create database connection")
	}

	database.DBConn.AutoMigrate(&domain.User{})
	database.DBConn.AutoMigrate(&domain.Product{})
}
