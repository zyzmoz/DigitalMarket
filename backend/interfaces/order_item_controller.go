package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zyzmoz/DigitalMarket/domain"
	"github.com/zyzmoz/DigitalMarket/usecases"
	"gorm.io/gorm"
)

type OrderItemController struct {
	OrderItemInteractor usecases.OrderItemInteractor
	Logger              usecases.Logger
}

func NewOrderItemController(db *gorm.DB, logger usecases.Logger) *OrderItemController {
	return &OrderItemController{
		OrderItemInteractor: usecases.OrderItemInteractor{
			OrderItemRepository: &OrderItemRepository{
				DB: db,
			},
		},
	}
}

// functions

func (c *OrderItemController) FindAll(ctx *fiber.Ctx) error {
	c.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	orderItems, err := c.OrderItemInteractor.FindAll()
	if err != nil {
		c.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(orderItems)
}

func (c *OrderItemController) FindOne(ctx *fiber.Ctx) error {
	c.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	id := ctx.Params("id")

	orderItem, err := c.OrderItemInteractor.FindOne(id)
	if err != nil {
		c.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(orderItem)
}

func (c *OrderItemController) Create(ctx *fiber.Ctx) error {
	c.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	var orderItemData domain.OrderItem
	if err := ctx.BodyParser(&orderItemData); err != nil {
		c.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)
		return err
	}

	orderItemData.ID = uuid.New().String()

	orderItem, err := c.OrderItemInteractor.Create(orderItemData)
	if err != nil {
		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(orderItem)
}

func (c *OrderItemController) Update(ctx *fiber.Ctx) error {
	c.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	var orderItemData domain.OrderItem
	if err := ctx.BodyParser(&orderItemData); err != nil {
		c.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)

		return err
	}

	orderItem, err := c.OrderItemInteractor.Update(orderItemData)
	if err != nil {
		c.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}
	return ctx.JSON(orderItem)
}

func (c *OrderItemController) Delete(ctx *fiber.Ctx) error {
	c.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	id := ctx.Params("id")

	orderItem, err := c.OrderItemInteractor.FindOne(id)
	if err != nil {
		c.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)

		return err
	}

	c.OrderItemInteractor.Delete(orderItem)

	return ctx.SendString("Payment Deleted")
}
