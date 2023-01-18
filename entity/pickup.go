package entity

import "gorm.io/gorm"

type Pickup struct {
	ID uint
	UserID uint
	ReservationsID uint
	PickupStatusID uint
	gorm.Model `json:"-"`
}