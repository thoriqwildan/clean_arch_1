package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/thoriqwildan/clean_arch_1/internal/model"
	"github.com/thoriqwildan/clean_arch_1/internal/usecase"
)

type ChannelController struct {
	UseCase *usecase.ChannelUseCase
	Log *logrus.Logger
}

func NewChannelController(useCase *usecase.ChannelUseCase, log *logrus.Logger) *ChannelController {
	return &ChannelController{
		UseCase: useCase,
		Log: log,
	}
}

func (cc *ChannelController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateChannelRequest)
	if err := ctx.BodyParser(request); err != nil {
		cc.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := cc.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		cc.Log.WithError(err).Error("Failed to create payment channel")
		return err
	}

	return ctx.JSON(model.WebResponse[model.ChannelResponse]{Data: *response})
}