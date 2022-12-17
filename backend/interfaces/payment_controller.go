package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zyzmoz/DigitalMarket/domain"
	"github.com/zyzmoz/DigitalMarket/usecases"
	"gorm.io/gorm"
)

type PaymentController struct {
	PaymentInteractor usecases.PaymentInteractor
	Logger            usecases.Logger
}

func NewPaymentController(db *gorm.DB, logger usecases.Logger) *PaymentController {
	return &PaymentController{
		PaymentInteractor: usecases.PaymentInteractor{
			PaymentRepository: &PaymentRepository{
				DB: db,
			},
		},
		Logger: logger,
	}
}

func (pc *PaymentController) FindAll(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	payment, err := pc.PaymentInteractor.FindAll()
	if err != nil {
		pc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(payment)
}

func (pc *PaymentController) FindOne(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	id := ctx.Params("id")

	payment, err := pc.PaymentInteractor.FindOne(id)
	if err != nil {
		pc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(payment)
}

func (pc *PaymentController) Create(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	var paymentData domain.Payment
	if err := ctx.BodyParser(&paymentData); err != nil {
		pc.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)
		return err
	}

	paymentData.ID = uuid.New().String()

	payment, err := pc.PaymentInteractor.Create(paymentData)
	if err != nil {
		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(payment)
}

func (pc *PaymentController) Update(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	var paymentData domain.Payment
	if err := ctx.BodyParser(&paymentData); err != nil {
		pc.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)

		return err
	}

	payment, err := pc.PaymentInteractor.Update(paymentData)
	if err != nil {
		pc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}
	return ctx.JSON(payment)
}

func (pc *PaymentController) Delete(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	id := ctx.Params("id")

	payment, err := pc.PaymentInteractor.FindOne(id)
	if err != nil {
		pc.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)

		return err
	}

	pc.PaymentInteractor.Delete(payment)

	return ctx.SendString("Payment Deleted")
}
