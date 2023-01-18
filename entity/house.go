package entity

import "gorm.io/gorm"

type House struct {
	ID uint
	Name string
	UserID uint 
	PricePerNight int 
	Description string
	CityID uint 
	MaxGuests int 
	gorm.Model `json:"-"`
}