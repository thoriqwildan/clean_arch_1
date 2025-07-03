package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/thoriqwildan/clean_arch_1/internal/entity"
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
