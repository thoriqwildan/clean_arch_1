package main

import (
	"strconv"

	"github.com/thoriqwildan/clean_arch_1/internal/config"
)

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

	webPort := ":" + strconv.Itoa(viperConfig.GetInt("web.port"))
	err := app.Listen(webPort)
	if err != nil {
		log.Fatal("Failed to start server: " + err.Error())
	}

	log.Info("Server started on port " + webPort)
}