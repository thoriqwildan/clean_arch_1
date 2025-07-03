package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thoriqwildan/clean_arch_1/internal/delivery/http"
)

type RouteConfig struct {
	App *fiber.App
	MethodController *http.MethodController
	ChannelController *http.ChannelController
}

func (rc *RouteConfig) Setup() {
	rc.SetupGeneralRoutes()
}

func (rc *RouteConfig) SetupGeneralRoutes() {
	rc.App.Post("/api/methods", rc.MethodController.Create)
	rc.App.Get("/api/methods", rc.MethodController.Filter)
	rc.App.Get("/api/methods/:id", rc.MethodController.FindById)
	rc.App.Put("/api/methods/:id", rc.MethodController.Update)
	rc.App.Delete("/api/methods/:id", rc.MethodController.Delete)

	rc.App.Post("/api/channels", rc.ChannelController.Create)
}