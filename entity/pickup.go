package entity

import (
	"time"

	"gorm.io/gorm"
)

type PickUp struct {
	ID             uint
	UserID         uint
	ReservationID  uint
	PickupStatusID uint
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

func (PickUp) TableName() string {
	return "pickups"
}
