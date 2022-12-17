package infra

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/zyzmoz/DigitalMarket/interfaces"
	"github.com/zyzmoz/DigitalMarket/usecases"
	"gorm.io/gorm"
)

func Dispatch(app *fiber.App, db *gorm.DB, logger usecases.Logger) {
	userCtrl := interfaces.NewUserController(db, logger)
	productCtrl := interfaces.NewProductController(db, logger)
	orderCtrl := interfaces.NewOrderController(db, logger)
	paymentCtrl := interfaces.NewPaymentController(db, logger)

	app.Get("/api/v1/user/:id", userCtrl.FindOne)
	app.Get("/api/v1/user", userCtrl.FindAll)
	app.Post("/api/v1/user", userCtrl.Create)
	app.Put("/api/v1/user", userCtrl.Update)
	app.Delete("/api/v1/user/:id", userCtrl.Delete)

	app.Get("/api/v1/product/:id", productCtrl.FindOne)
	app.Get("/api/v1/product", productCtrl.FindAll)
	app.Post("/api/v1/product", productCtrl.Create)
	app.Put("/api/v1/product", productCtrl.Update)
	app.Delete("/api/v1/product/:id", productCtrl.Delete)

	app.Get("/api/v1/order/:id", orderCtrl.FindOne)
	app.Get("/api/v1/order", orderCtrl.FindAll)
	app.Post("/api/v1/order", orderCtrl.Create)
	app.Put("/api/v1/order", orderCtrl.Update)
	app.Delete("/api/v1/order/:id", orderCtrl.Delete)

	app.Get("/api/v1/payment/:id", paymentCtrl.FindOne)
	app.Get("/api/v1/payment", paymentCtrl.FindAll)
	app.Post("/api/v1/payment", paymentCtrl.Create)
	app.Put("/api/v1/payment", paymentCtrl.Update)
	app.Delete("/api/v1/payment/:id", paymentCtrl.Delete)

	app.Listen(":" + os.Getenv("SERVER_PORT"))
}
