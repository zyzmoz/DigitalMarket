package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/zyzmoz/DigitalMarket/domain"
	"github.com/zyzmoz/DigitalMarket/usecases"
)

type ProductController struct {
	ProductInteractor usecases.ProductInteractor
	Logger            usecases.Logger
}

func NewProductController(db *gorm.DB, logger usecases.Logger) *ProductController {
	return &ProductController{
		ProductInteractor: usecases.ProductInteractor{
			ProductRepository: &ProductRepository{
				DB: db,
			},
		},
		Logger: logger,
	}
}

func (pc *ProductController) FindAll(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	products, err := pc.ProductInteractor.FindAll()
	if err != nil {
		pc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(products)
}

func (pc *ProductController) FindOne(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	id := ctx.Params("id")

	product, err := pc.ProductInteractor.FindOne(id)
	if err != nil {
		pc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(product)
}

func (pc *ProductController) Create(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())
	var productData domain.Product
	if err := ctx.BodyParser(&productData); err != nil {
		pc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	productData.ID = uuid.New().String()

	product, err := pc.ProductInteractor.Create(productData)
	if err != nil {
		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(product)
}

func (pc *ProductController) Update(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	var productData domain.Product
	if err := ctx.BodyParser(&productData); err != nil {
		pc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	product, err := pc.ProductInteractor.Update(productData)
	if err != nil {
		pc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(product)
}

func (pc *ProductController) Delete(ctx *fiber.Ctx) error {
	pc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())
	id := ctx.Params("id")

	product, err := pc.ProductInteractor.FindOne(id)
	if err != nil {
		pc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	pc.ProductInteractor.Delete(product)

	return ctx.SendString("Product Deleted")
}
