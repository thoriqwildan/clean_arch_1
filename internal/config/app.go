package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoriqwildan/clean_arch_1/internal/delivery/http"
	"github.com/thoriqwildan/clean_arch_1/internal/delivery/http/route"
	"github.com/thoriqwildan/clean_arch_1/internal/repository"
	"github.com/thoriqwildan/clean_arch_1/internal/usecase"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB *gorm.DB
	App *fiber.App
	Log *logrus.Logger
	Validate *validator.Validate
	config *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	methodRepository := repository.NewMethodRepository(config.Log)

	methodUseCase := usecase.NewMethodUseCase(config.DB, config.Log, config.Validate, methodRepository)

	methodController := http.NewMethodController(methodUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App: config.App,
		MethodController: methodController,
	}

	routeConfig.Setup()
}