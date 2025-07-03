package usecase

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/thoriqwildan/clean_arch_1/internal/entity"
	"github.com/thoriqwildan/clean_arch_1/internal/model"
	"github.com/thoriqwildan/clean_arch_1/internal/model/converter"
	"github.com/thoriqwildan/clean_arch_1/internal/repository"
	"gorm.io/gorm"
)

type ChannelUseCase struct {
	DB *gorm.DB
	Log *logrus.Logger
	ChannelRepository *repository.ChannelRepository
	Validate *validator.Validate
}

func NewChannelUseCase(db *gorm.DB, log *logrus.Logger, channelRepository *repository.ChannelRepository, validate *validator.Validate) *ChannelUseCase {
	return &ChannelUseCase{
		DB: db,
		Log: log,
		ChannelRepository: channelRepository,
		Validate: validate,
	}
}

func (c *ChannelUseCase) Create(ctx context.Context, request *model.CreateChannelRequest) (*model.ChannelResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation Error")
		return nil, fiber.ErrBadRequest
	}

	if err := c.ChannelRepository.FindChannelByName(tx, request.Name); err == nil {
    c.Log.WithField("name", request.Name).Error("Payment channel with this name already exists")
    return nil, fiber.ErrConflict
  } else if !errors.Is(err, gorm.ErrRecordNotFound) {
      // Tangani error lain selain 'record not found'
      c.Log.WithError(err).Error("Failed to check for existing payment channel by name")
      return nil, err
  }

	if err := c.ChannelRepository.FindChannelByCode(tx, request.Code); err == nil {
    c.Log.WithField("code", request.Code).Error("Payment channel with this code already exists")
    return nil, fiber.ErrConflict   
  } else if !errors.Is(err, gorm.ErrRecordNotFound) {
      // Tangani error lain selain 'record not found'
      c.Log.WithError(err).Error("Failed to check for existing payment channel by code")
      return nil, err
  }

	if err := c.ChannelRepository.FindMethodById(tx, request.PaymentMethodId); err != nil {
    c.Log.WithError(err).Error("Payment method not found")
    return nil, fiber.ErrNotFound
  }


	channel := &entity.PaymentChannel{
		PaymentMethodId: request.PaymentMethodId,
		Code: request.Code,
		Name: request.Name,
		IconUrl: sql.NullString{String: request.IconUrl, Valid: request.IconUrl != ""},
		OrderNum: sql.NullInt64{Int64: int64(request.OrderNum), Valid: request.OrderNum > 0},
		LibName: sql.NullString{String: request.LibName, Valid: request.LibName != ""},
		UserAction: request.UserAction,
		MDR: strconv.Itoa(request.Mdr),
		FixedFee: float64(request.FixedFee),
	}

	if err := c.ChannelRepository.Create(tx, channel); err != nil {
		c.Log.WithError(err).Error("Failed to create payment channel")
		return nil, err
	}

	method := &entity.PaymentMethod{}
	if err := c.ChannelRepository.FindMethodByChannel(tx, method, request.PaymentMethodId); err != nil {
		c.Log.WithError(err).Error("Failed to find payment method")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, err
	}

	return converter.ChannelToResponse(channel, method), nil
}