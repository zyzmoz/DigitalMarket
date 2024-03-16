package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zyzmoz/DigitalMarket/database"
	"github.com/zyzmoz/DigitalMarket/infra"
)

func main() {
	app := fiber.New()

	logger := infra.NewLogger()

	infra.Load(logger)

	host := os.Getenv("HOST")

	app.Use(cors.New(cors.Config{
		AllowOrigins: host,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	infra.InitDb()

	infra.Dispatch(app, database.DBConn, logger)
}
