package entity

import (
	"time"

	"gorm.io/gorm"
)

type HousePhoto struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	HouseID   uint      `json:"house_id"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (a *HousePhoto) TableName() string {
	return "houses_photos"
}
