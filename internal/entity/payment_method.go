package entity

import (
	"database/sql"
	"time"
)

type PaymentMethod struct {
	ID 								uint
	Name 							string	`gorm:"unique"`
	Desc 			sql.NullString
	OrderNum					int	`gorm:"default:1"`
	UserAction				string
	CreatedAt					time.Time
	UpdatedAt					time.Time
	Code 							sql.NullString

	Channels	[]PaymentChannel `gorm:"foreignKey:PaymentMethodId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}