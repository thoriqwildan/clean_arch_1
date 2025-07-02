package entity

import (
	"database/sql"
	"time"
)

type PaymentChannel struct {
	ID 				uint
	PaymentMethodId	uint
	Code 			string	`gorm:"unique"`
	Name 			string `gorm:"unique"`
	IconUrl		sql.NullString
	OrderNum	sql.NullInt64	`gorm:"default:1"`
	LibName		sql.NullString
	UserAction	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	MDR       string  `gorm:"default:'0'"`
	FixedFee  float64 `gorm:"default:0"`
}