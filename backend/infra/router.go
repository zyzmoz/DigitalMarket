package infra

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zyzmoz/DigitalMarket/interfaces"
	"github.com/zyzmoz/DigitalMarket/usecases"
)

func Dispatch(app *fiber.App, db *gorm.DB, logger usecases.Logger) {
	userCtrl := interfaces.NewUserController(db, logger)

	app.Get("/api/v1/user/:id", userCtrl.FindOne)
	app.Get("/api/v1/user", userCtrl.FindAll)
	app.Post("/api/v1/user", userCtrl.Create)
	app.Put("/api/v1/user", userCtrl.Update)
	app.Delete("/api/v1/user/:id", userCtrl.Delete)

	app.Listen(":" + os.Getenv("SERVER_PORT"))
	defer db.Close()
}
