package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zyzmoz/DigitalMarket/domain"
	"github.com/zyzmoz/DigitalMarket/usecases"
	"gorm.io/gorm"
)

type OrderController struct {
	OrderInteractor usecases.OrderInteractor
	Logger          usecases.Logger
}

func NewOrderController(db *gorm.DB, logger usecases.Logger) *OrderController {
	return &OrderController{
		OrderInteractor: usecases.OrderInteractor{
			OrderRepository: &OrderRepository{
				DB: db,
			},
		},
		Logger: logger,
	}
}

func (oc *OrderController) FindAll(ctx *fiber.Ctx) error {
	oc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	orders, err := oc.OrderInteractor.FindAll()
	if err != nil {
		oc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(orders)
}

func (oc *OrderController) FindOne(ctx *fiber.Ctx) error {
	oc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	id := ctx.Params("id")

	order, err := oc.OrderInteractor.FindOne(id)
	if err != nil {
		oc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(order)
}

func (oc *OrderController) Create(ctx *fiber.Ctx) error {
	oc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	var orderData domain.Order
	if err := ctx.BodyParser(&orderData); err != nil {
		oc.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)
		return err
	}

	orderData.ID = uuid.New().String()

	order, err := oc.OrderInteractor.Create(orderData)
	if err != nil {
		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(order)
}

func (oc *OrderController) Update(ctx *fiber.Ctx) error {
	oc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	var orderData domain.Order
	if err := ctx.BodyParser(&orderData); err != nil {
		oc.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)

		return err
	}

	order, err := oc.OrderInteractor.Update(orderData)
	if err != nil {
		oc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(order)
}

func (oc *OrderController) Delete(ctx *fiber.Ctx) error {
	oc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	id := ctx.Params("id")

	order, err := oc.OrderInteractor.FindOne(id)
	if err != nil {
		oc.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)

		return err
	}

	oc.OrderInteractor.Delete(order)

	return ctx.SendString("Order Deleted")
}
