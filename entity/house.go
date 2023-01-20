package entity

import "gorm.io/gorm"

type House struct {
	ID uint `json:"id" gorm:"primaryKey" `
	Name string `json:"name" validate:"required"`
	UserID uint `json:"user_id" validate:"required"`
	PricePerNight int `json:"price_per_night" validate:"required"`
	Description string `json:"description" validate:"required"`
	CityID uint `json:"city_id" validate:"required"`
	MaxGuests int `json:"max_guests"`
	gorm.Model `json:"-"`
}