package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	ID uint
	HouseID uint
	UserID uint
	CheckInDate time.Time
	CheckOutDate time.Time
	TotalPrice float64
	gorm.Model `json:"-"`
}