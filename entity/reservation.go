package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	HouseID      uint           `json:"house_id"`
	UserID       uint           `json:"user_id"`
	CheckInDate  time.Time      `json:"check_in"`
	CheckOutDate time.Time      `json:"check_out"`
	TotalPrice   float64        `json:"total_price"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
