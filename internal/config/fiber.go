package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(viper *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Clean Architecture",
		Prefork: viper.GetBool("web.prefork"),
	})

	return app
}