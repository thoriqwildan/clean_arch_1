package main

import (
	"strconv"

	"github.com/gofiber/swagger"

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

	app.Get("/swagger/*", swagger.HandlerDefault)

	webPort := ":" + strconv.Itoa(viperConfig.GetInt("web.port"))
	err := app.Listen(webPort)
	if err != nil {
		log.Fatal("Failed to start server: " + err.Error())
	}

	log.Info("Server started on port " + webPort)
}