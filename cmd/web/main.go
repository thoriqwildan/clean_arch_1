package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	_ "github.com/thoriqwildan/clean_arch_1/docs"
	"github.com/thoriqwildan/clean_arch_1/internal/config"
)

// @title SVD Clone API
// @version 1.0
// @description This is a sample swagger for Fiber
// @host localhost:3000
// @BasePath /
func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	app := config.NewFiber(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validator := config.NewValidator(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB: db,
		App: app,
		Log: log,
		Validate: validator,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/reference", func(ctx *fiber.Ctx) error {
		html, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			DarkMode: true,
		})
		if err != nil {
			log.Error("Failed to generate API reference HTML: " + err.Error())
			return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to generate API reference HTML")
		}
		return ctx.Type("html").SendString(html)
	})

	webPort := ":" + strconv.Itoa(viperConfig.GetInt("web.port"))
	err := app.Listen(webPort)
	if err != nil {
		log.Fatal("Failed to start server: " + err.Error())
	}

	log.Info("Server started on port " + webPort)
}