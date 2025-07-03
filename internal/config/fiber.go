package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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
			return ctx.Status(e.Code).JSON(fiber.Map{
				"error": e.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
}