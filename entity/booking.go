package entity

import "gorm.io/gorm"

type Booking struct {
	ID uint 
	UserID uint 
	HouseID uint 
	ReservationID uint 
	gorm.Model 
}