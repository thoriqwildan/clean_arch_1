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

	return ctx.JSON(model.WebResponse[model.ChannelResponse]{Success: true, Data: *response})
}

func (cc *ChannelController) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		cc.Log.Error("ID parameter is required")
		return fiber.ErrBadRequest
	}

	response, err := cc.UseCase.GetChannelById(ctx.UserContext(), id)
	if err != nil {
		cc.Log.WithError(err).Error("Failed to find payment channel")
		return err
	}
	return ctx.JSON(model.WebResponse[model.ChannelResponse]{Success: true, Data: *response})
}

func (cc *ChannelController) Get(ctx *fiber.Ctx) error {
	query := &model.FilterChannelQuery{
		Code: ctx.Query("code"),
		Name: ctx.Query("name"),
		Page: ctx.QueryInt("page", 1),
		Limit: ctx.QueryInt("limit", 10),
	}

	responses, total, err := cc.UseCase.Get(ctx.UserContext(), query)
	if err != nil {
		cc.Log.WithError(err).Error("Failed to get payment channels")
		return err
	}

	paging := &model.PaginationPage{
		Page: query.Page,
		Limit: query.Limit,
		Total: total,
	}

	return ctx.JSON(model.WebResponse[[]model.ChannelResponse]{
		Success: true,
		Data: responses,
		Meta: paging,
	})
}

func (cc *ChannelController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		cc.Log.Error("ID parameter is required")
		return fiber.ErrBadRequest
	}

	request := new(model.UpdateChannelRequest)
	if err := ctx.BodyParser(request); err != nil {
		cc.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.ID = id
	response, err := cc.UseCase.UpdateChannel(ctx.UserContext(), request)
	if err != nil {
		cc.Log.WithError(err).Error("Failed to update payment channel")
		return err
	}

	return ctx.JSON(model.WebResponse[model.ChannelResponse]{Success: true, Data: *response})
}

func (cc *ChannelController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		cc.Log.Error("ID parameter is required")
		return fiber.ErrBadRequest
	}

	if err := cc.UseCase.Delete(ctx.UserContext(), id); err != nil {
		cc.Log.WithError(err).Error("Failed to delete payment channel")
		return err
	}

	return ctx.JSON(model.WebResponse[any]{Success: true, Data: nil})
}