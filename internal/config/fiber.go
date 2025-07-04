package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/thoriqwildan/clean_arch_1/internal/helper"
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
				Error:   err.Error(),
			})
		} else if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorsMap := helper.TranslateErrorMessage(validationErrors)
			return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
				Success: false,
				Error:   errorsMap,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Success: false,
			Error:   "Internal Server Error",
		})
	}
}