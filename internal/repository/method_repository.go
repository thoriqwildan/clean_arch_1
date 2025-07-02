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


func (mr *MethodRepository) FilterMethod(request *model.FilterMethodRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if code := request.Code; code != "" {
			code = "%" + code + "%"
			tx = tx.Where("code LIKE ?", code)
		}
		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}
	return tx
	}
}

func (mr *MethodRepository) Search(db *gorm.DB, request *model.FilterMethodRequest) ([]entity.PaymentMethod, int64, error) {
	var methods []entity.PaymentMethod
	if err := db.Scopes(mr.FilterMethod(request)).Offset((request.Page - 1) * request.Limit).Find(&methods).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.PaymentMethod{}).Scopes(mr.FilterMethod(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return methods, total, nil
}




