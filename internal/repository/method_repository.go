package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/thoriqwildan/clean_arch_1/internal/entity"
	"github.com/thoriqwildan/clean_arch_1/internal/model"
	"gorm.io/gorm"
)

type MethodRepository struct {
	Repository[entity.PaymentMethod]
	Log *logrus.Logger
}

func NewMethodRepository(log *logrus.Logger) *MethodRepository {
	return &MethodRepository{
		Log: log,
	}
}

func (mr *MethodRepository) FindById(db *gorm.DB, method *entity.PaymentMethod, id any) error {
	return db.Where("id = ?", id).Take(method).Error
}

func (mr *MethodRepository) FilterMethod(db *gorm.DB, request *model.FilterMethodRequest) ([]entity.PaymentMethod, int64, error) {
	var methods []entity.PaymentMethod
	var total int64

	query := db.Model(&entity.PaymentMethod{})

	if request.Code != "" {
		query = query.Where("code ILIKE ?", "%"+request.Code+"%")
	}
	if request.Name != "" {
		query = query.Where("name ILIKE ?", "%"+request.Name+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (request.Page - 1) * request.Limit

	if request.Page < 1 {
		request.Page = 1
	}
	if request.Limit < 1 {
		request.Limit = 10
	}

	if err := query.Limit(request.Limit).Offset(offset).Order("id DESC").Find(&methods).Error; err != nil {
		mr.Log.WithError(err).Error("Failed to filter payment methods")
		return nil, 0, err
	}

	return methods, total, nil
}