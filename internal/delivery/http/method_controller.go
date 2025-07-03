package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/thoriqwildan/clean_arch_1/internal/model"
	"github.com/thoriqwildan/clean_arch_1/internal/usecase"
)

type MethodController struct {
	UseCase *usecase.MethodUseCase
	Log *logrus.Logger
}

func NewMethodController(useCase *usecase.MethodUseCase, log *logrus.Logger) *MethodController {
	return &MethodController{
		UseCase: useCase,
		Log: log,
	}
}

// @Router /api/methods [post]
// @Param  request body model.CreateMethodRequest true "Create Method Request"
// @Success 200 {object} model.MethodResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Methods
// @Summary Create a new payment method
// @Description Create a new payment method with the provided details.
// @Accept json
func (mc *MethodController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateMethodRequest)
	if err := ctx.BodyParser(request); err != nil {
		mc.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := mc.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		mc.Log.WithError(err).Error("Failed to create payment method")
		return err
	}

	return ctx.JSON(model.WebResponse[model.MethodResponse]{Success: true, Data: *response})
}

// @Router /api/methods/{id} [get]
// @Success 200 {object} model.MethodResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Methods
// @Summary Find a payment method by ID
// @Description Find a payment method by its ID.
// @Param id path string true "Method ID"
// @Accept json
func (mc *MethodController) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		mc.Log.Error("ID parameter is required")
		return fiber.ErrBadRequest
	}

	method, err := mc.UseCase.GetMethodById(ctx.UserContext(), id)
	if err != nil {
		mc.Log.WithError(err).Error("Failed to find payment method")
		return err
	}
	return ctx.JSON(model.WebResponse[model.MethodResponse]{Success: true, Data: *method})
}

// @Router /api/methods [get]
// @Success 200 {object} model.MethodsResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Methods
// @Summary Find a payment methods
// @Description Find a payment methods by its ID.
// @Param   name query     string            false       "Filter by method name"
// @Param   code query     string            false       "Filter by method code"
// @Param   page query     int               false       "Page number" default(1)
// @Param   limit query    int               false       "Number of items per page" default
// @Accept json
func (mc *MethodController) Filter(ctx *fiber.Ctx) error {
	request := &model.FilterMethodRequest{
		Code: ctx.Query("code"),
		Name: ctx.Query("name"),
		Page: ctx.QueryInt("page", 1),
		Limit: ctx.QueryInt("limit", 10),
	}

	responses, total, err := mc.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		mc.Log.WithError(err).Error("Failed to filter payment methods")
		return err
	}

	paging := &model.PaginationPage{
		Page: request.Page,
		Limit: request.Limit,
		Total: total,
	}

	return ctx.JSON(model.WebResponse[[]model.MethodResponse]{
		Success: true,
		Data: responses,
		Meta: paging,
	})
}

// @Router /api/methods/{id} [put]
// @Param  request body model.UpdateMethodRequest true "Update Method Request"
// @Param id path string true "Method ID"
// @Success 200 {object} model.MethodResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Methods
// @Summary Update a payment method
// @Description Update a payment method with the provided details.
// @Accept json
func (mc *MethodController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateMethodRequest)
	if err := ctx.BodyParser(request); err != nil {
		mc.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("id")

	response, err := mc.UseCase.UpdateMethod(ctx.UserContext(), request)
	if err != nil {
		mc.Log.WithError(err).Error("Failed to update payment method")
		return err
	}

	return ctx.JSON(model.WebResponse[model.MethodResponse]{Success: true, Data: *response})
}

// @Router /api/methods/{id} [delete]
// @Param id path string true "Method ID"
// @Success 200 {object} model.MethodResponseWrapper
// @Failure 400 {object} model.ErrorWrapper "Bad Request"
// @Failure 500 {object} model.ErrorWrapper "Internal Server Error"
// @Tags Methods
// @Summary Delete a payment method
// @Description Delete a payment method with the provided details.
// @Accept json
func (mc *MethodController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		mc.Log.Error("ID parameter is required")
		return fiber.ErrBadRequest
	}

	if err := mc.UseCase.DeleteMethod(ctx.UserContext(), id); err != nil {
		mc.Log.WithError(err).Error("Failed to delete payment method")
		return err
	}

	return ctx.JSON(model.WebResponse[any]{Success: true, Data: nil})
}