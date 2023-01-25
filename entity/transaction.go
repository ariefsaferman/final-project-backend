package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint    `json:"id"`
	SourceOfFundsId uint    `json:"source_of_funds_id"`
	UserID          uint    `json:"user_id"`
	Balance         float64 `json:"balance"`
	TypeTransaction SourceOfFund `json:"type_transaction" gorm:"foreignKey:SourceOfFundsId"`
	gorm.Model      `json:"-"`
}
