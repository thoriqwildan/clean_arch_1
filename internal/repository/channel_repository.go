package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/thoriqwildan/clean_arch_1/internal/entity"
	"github.com/thoriqwildan/clean_arch_1/internal/model"
	"gorm.io/gorm"
)

type ChannelRepository struct {
	Repository[entity.PaymentChannel]
	Log *logrus.Logger
}

func NewChannelRepository(log *logrus.Logger) *ChannelRepository {
	return &ChannelRepository{
		Log: log,
	}
}

func (cr *ChannelRepository) FindChannelByCode(db *gorm.DB, code string) error {
	return db.Where("code = ?", code).Take(&entity.PaymentChannel{}).Error
}

func (cr *ChannelRepository) FindChannelByName(db *gorm.DB, name string) error {
	return db.Where("name = ?", name).Take(&entity.PaymentChannel{}).Error
} 

func (cr *ChannelRepository) FindMethodByChannel(db *gorm.DB, method *entity.PaymentMethod, id any) error {
	return db.Where("id = ?", id).Take(method).Error
}

func (cr *ChannelRepository) FindMethodById(db *gorm.DB, id any) error {
	return db.Where("id = ?", id).Take(&entity.PaymentMethod{}).Error
}


func (cr *ChannelRepository) FilterChannel(request *model.FilterChannelQuery) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if code := request.Code; code != "" {
			code = "%" + code + "%"
			tx = tx.Where("code ILIKE ?", code)
		}
		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name ILIKE ?", name)
		}
		return tx
	}
}

func (cr *ChannelRepository) Search(db *gorm.DB, request *model.FilterChannelQuery) ([]entity.PaymentChannel, int64, error) {
	var channels []entity.PaymentChannel
	var total int64

	baseQuery := db.Model(&entity.PaymentChannel{}).Scopes(cr.FilterChannel(request))

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := baseQuery.
		Offset((request.Page - 1) * request.Limit).
		Limit(request.Limit).
		Find(&channels).Error
	if err != nil {
		return nil, 0, err
	}

	return channels, total, nil
}