package usecase

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/thoriqwildan/clean_arch_1/internal/entity"
	"github.com/thoriqwildan/clean_arch_1/internal/model"
	"github.com/thoriqwildan/clean_arch_1/internal/model/converter"
	"github.com/thoriqwildan/clean_arch_1/internal/repository"
	"gorm.io/gorm"
)

type MethodUseCase struct {
	DB  *gorm.DB
	Log *logrus.Logger
	Validate *validator.Validate
	MethodRepository *repository.MethodRepository
}

func NewMethodUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, methodRepository *repository.MethodRepository) *MethodUseCase {
	return &MethodUseCase{
		DB: db,
		Log: log,
		Validate: validate,
		MethodRepository: methodRepository,
	}
} 

func (m *MethodUseCase) Create(ctx context.Context, request *model.CreateMethodRequest) (*model.MethodResponse, error) {
	tx := m.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := m.Validate.Struct(request); err != nil {
		m.Log.WithError(err).Error("Validation error")
		return nil, fiber.ErrBadRequest
	}

	if err := m.MethodRepository.FindMethodByName(tx, request.Name); err == nil {
		m.Log.WithField("name", request.Name).Error("Payment method with this name already exists")
		return nil, fiber.ErrConflict
	}

	method := &entity.PaymentMethod{
		Name: request.Name,
		Desc: sql.NullString{String: request.Desc, Valid: request.Desc != ""},
		OrderNum: request.OrderNum,
		UserAction: request.UserAction,
		Code: sql.NullString{String: request.Code, Valid: request.Code != ""},
	}

	if err := m.MethodRepository.Create(tx, method); err != nil {
		m.Log.WithError(err).Error("Failed to create payment method")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		m.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.MethodToResponse(method), nil
}

func (m *MethodUseCase) GetMethodById(ctx context.Context, id any) (*model.MethodResponse, error) {
	tx := m.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	method := &entity.PaymentMethod{}
	if err := m.MethodRepository.FindById(tx, method, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			m.Log.WithError(err).Error("Payment method not found")
			return nil, fiber.ErrNotFound
		}
		m.Log.WithError(err).Error("Failed to find payment method")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		m.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.MethodToResponse(method), nil
}

func (m *MethodUseCase) Get(ctx context.Context, request *model.FilterMethodRequest) ([]model.MethodResponse, int64, error) {
	tx := m.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := m.Validate.Struct(request); err != nil {
		m.Log.WithError(err).Error("Validation error")
		return nil, 0, fiber.ErrBadRequest
	}

	methods, total, err := m.MethodRepository.Search(tx, request)
	if err != nil {
		m.Log.WithError(err).Error("Failed to filter payment methods")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		m.Log.WithError(err).Error("Failed to commit transaction")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.MethodResponse, len(methods))
	for i, method := range methods {
		responses[i] = *converter.MethodToResponse(&method)
	}

	return responses, total, nil
}

func (m *MethodUseCase) UpdateMethod(ctx context.Context, request *model.UpdateMethodRequest) (*model.MethodResponse, error) {
	tx := m.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	method := &entity.PaymentMethod{}
	if err := m.MethodRepository.FindById(tx, method, request.ID); err != nil {
		m.Log.WithError(err).Error("Payment method not found")
		return nil, fiber.ErrNotFound
	}

	if err := m.Validate.Struct(request); err != nil {
		m.Log.WithError(err).Error("Validation error")
		return nil, fiber.ErrBadRequest
	}

	method.Name = request.Name
	method.Desc = sql.NullString{String: request.Desc, Valid: request.Desc != ""}
	method.OrderNum = request.OrderNum
	method.UserAction = request.UserAction
	method.Code = sql.NullString{String: request.Code, Valid: request.Code != ""}

	if err := m.MethodRepository.Update(tx, method); err != nil {
		m.Log.WithError(err).Error("Failed to update payment method")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		m.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.MethodToResponse(method), nil
}

func (m *MethodUseCase) DeleteMethod(ctx context.Context, id any) error {
	tx := m.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	method := new(entity.PaymentMethod)
	if err := m.MethodRepository.FindById(tx, method, id); err != nil {
		m.Log.WithError(err).Error("Payment method not found")
		return fiber.ErrNotFound
	}

	if err := m.MethodRepository.Delete(tx, method); err != nil {
		m.Log.WithError(err).Error("Failed to delete payment method")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		m.Log.WithError(err).Error("Failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}