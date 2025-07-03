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

// @Router /api/channels [post]
// @Param  request body model.CreateChannelRequest true "Create Channel Request"
// @Success 200 {object} model.ChannelResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Channels
// @Summary Create a new payment channel
// @Description Create a new payment channel with the provided details.
// @Accept json
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

// @Router /api/channels/{id} [get]
// @Success 200 {object} model.ChannelResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Channels
// @Summary Find a payment channel by ID
// @Description Find a payment channel by its ID.
// @Param id path string true "Channel ID"
// @Accept json
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

// @Router /api/channels [get]
// @Success 200 {object} model.ChannelsResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Channels
// @Summary Find a payment channels
// @Description Find a payment channel by its ID.
// @Param   name query     string            false       "Filter by channel name"
// @Param   code query     string            false       "Filter by channel code"
// @Param   page query     int               false       "Page number" default(1)
// @Param   limit query    int               false       "Number of items per page" default
// @Accept json
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

// @Router /api/channels/{id} [put]
// @Param  request body model.UpdateChannelRequest true "Update Channel Request"
// @Param id path string true "Channel ID"
// @Success 200 {object} model.ChannelResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Channels
// @Summary Update a payment channel
// @Description Update a payment channel with the provided details.
// @Accept json
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

// @Router /api/channels/{id} [delete]
// @Param id path string true "Channel ID"
// @Success 200 {object} model.ChannelResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Channels
// @Summary Delete a payment channel
// @Description Delete a payment channel with the provided details.
// @Accept json
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