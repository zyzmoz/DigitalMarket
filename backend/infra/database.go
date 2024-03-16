package infra

import (
	"fmt"
	"os"

	"github.com/zyzmoz/DigitalMarket/database"
	"github.com/zyzmoz/DigitalMarket/domain"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDb() {
	var err error

	switch os.Getenv("DATABASE_DRIVER") {
	case "mssql":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))
		database.DBConn, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))
		database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		database.DBConn, err = gorm.Open(sqlite.Open("temp.db"), &gorm.Config{})
	}

	if err != nil {
		panic("failed to create database connection")
	}

	database.DBConn.AutoMigrate(&domain.User{})
	database.DBConn.AutoMigrate(&domain.Product{})
	database.DBConn.AutoMigrate(&domain.Order{})
	database.DBConn.AutoMigrate(&domain.Payment{})
}
