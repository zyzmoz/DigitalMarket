package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/zyzmoz/DigitalMarket/domain"
	"github.com/zyzmoz/DigitalMarket/usecases"
)

type UserController struct {
	UserInteractor usecases.UserInteractor
	Logger         usecases.Logger
}

func NewUserController(db *gorm.DB, logger usecases.Logger) *UserController {
	return &UserController{
		UserInteractor: usecases.UserInteractor{
			UserRepository: &UserRepository{
				DB: db,
			},
		},
		Logger: logger,
	}
}

func (uc *UserController) FindAll(ctx *fiber.Ctx) error {
	uc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	users, err := uc.UserInteractor.FindAll()
	if err != nil {
		uc.Logger.LogError("%s", err)

		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(users)
}

func (uc *UserController) FindOne(ctx *fiber.Ctx) error {
	uc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	id := ctx.Params("id")

	user, err := uc.UserInteractor.FindOne(id)
	if err != nil {
		uc.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(user)
}

func (uc *UserController) Create(ctx *fiber.Ctx) error {
	uc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	var userData domain.User
	if err := ctx.BodyParser(&userData); err != nil {
		ctx.Status(500).JSON(err)
		return err
	}

	user, err := uc.UserInteractor.Create(userData)
	if err != nil {
		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(user)
}

func (uc *UserController) Update(ctx *fiber.Ctx) error {
	uc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	var userData domain.User
	if err := ctx.BodyParser(&userData); err != nil {
		ctx.Status(500).JSON(err)
		return err
	}

	user, err := uc.UserInteractor.Update(userData)
	if err != nil {
		ctx.Status(500).JSON(err)
		return err
	}

	return ctx.JSON(user)
}

func (uc *UserController) Delete(ctx *fiber.Ctx) error {
	uc.Logger.LogAccess("%s %s %s\n", ctx.Context().RemoteAddr(), ctx.Context().Method(), ctx.Context().URI())

	id := ctx.Params("id")

	user, err := uc.UserInteractor.FindOne(id)
	if err != nil {
		uc.Logger.LogError("%s", err)
		ctx.Status(500).JSON(err)
		return err
	}

	uc.UserInteractor.Delete(user)

	return ctx.SendString("User Deleted")
}
