package entity

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type House struct {
	ID            uint       `json:"id" gorm:"primaryKey" `
	Name          string     `json:"name" validate:"required"`
	UserID        uint       `json:"user_id" validate:"required"`
	PricePerNight float64    `json:"price_per_night" validate:"required"`
	Description   string     `json:"description" validate:"required"`
	CityID        uint       `json:"city_id" validate:"required"`
	MaxGuests     uint       `json:"max_guests" gorm:"column:max_guest"`
	HousePhoto    HousePhoto `form:"house_photo"`
	gorm.Model    `json:"-"`
}

type HousePhotoRequest struct {
	PhotoUrl *multipart.FileHeader `form:"photo_url" validate:"required"`
}
