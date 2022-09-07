package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zyzmoz/DigitalMarket/database"
	"github.com/zyzmoz/DigitalMarket/infra"
)

func main() {
	app := fiber.New()
	logger := infra.NewLogger()

	infra.Load(logger)

	infra.InitDb()

	infra.Dispatch(app, database.DBConn, logger)
}
