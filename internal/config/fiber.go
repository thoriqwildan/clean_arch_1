package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/thoriqwildan/clean_arch_1/internal/model"
)

func NewFiber(viper *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Clean Architecture",
		Prefork: viper.GetBool("web.prefork"),
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		if e, ok := err.(*fiber.Error); ok {
			return ctx.Status(e.Code).JSON(&model.WebResponse[any]{
				Success: false,
				Error:   e.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Success: false,
			Error:   "Internal Server Error",
		})
	}
}