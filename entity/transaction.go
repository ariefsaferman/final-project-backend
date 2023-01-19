package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint
	SourceOfFundsId uint
	UserID          uint
	Balance         float64
	gorm.Model
}
